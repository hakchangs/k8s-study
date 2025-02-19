package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

func mutatePods(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusInternalServerError)
		return
	}

	// 요청을 admission review로 변환
	var admissionReviewReq admissionv1.AdmissionReview
	err = json.Unmarshal(body, &admissionReviewReq)
	if err != nil {
		http.Error(w, "could not unmarshal request", http.StatusInternalServerError)
		return
	}

	// Pod인지 확인
	var pod corev1.Pod
	err = json.Unmarshal(admissionReviewReq.Request.Object.Raw, &pod)
	if err != nil {
		http.Error(w, "could not unmarshal pod object", http.StatusInternalServerError)
		return
	}

	// 환경 변수 추가
	envVar := corev1.EnvVar{Name: "FOO", Value: "BAR"}
	pod.Spec.Containers[0].Env = append(pod.Spec.Containers[0].Env, envVar)

	// 패치 JSON 생성
	patch := `[{"op": "add", "path": "/spec/containers/0/env/-", "value": {"name": "FOO", "value": "BAR"}}]`
	patchType := admissionv1.PatchTypeJSONPatch

	// 응답 생성
	admissionResponse := admissionv1.AdmissionResponse{
		Allowed:   true,
		Patch:     []byte(patch),
		PatchType: &patchType,
		UID:       admissionReviewReq.Request.UID,
	}

	admissionReviewResp := admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
		Response: &admissionResponse,
	}

	respBytes, err := json.Marshal(admissionReviewResp)
	if err != nil {
		http.Error(w, "could not marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func main() {
	http.HandleFunc("/mutate", mutatePods)
	http.HandleFunc("/health", health)
	fmt.Println("Webhook server running on :8080")
	//http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8080", "/etc/certs/tls.crt", "/etc/certs/tls.key", nil)
}

func health(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("ok"))
}
