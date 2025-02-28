package example.spi;

import org.keycloak.authentication.AuthenticationFlowContext;
import org.keycloak.authentication.AuthenticationFlowError;
import org.keycloak.authentication.Authenticator;
import org.keycloak.models.KeycloakSession;
import org.keycloak.models.RealmModel;
import org.keycloak.models.UserModel;
import org.keycloak.sessions.AuthenticationSessionModel;

import java.util.Random;
import java.util.logging.Logger;

public class SmsOtpAuthenticator implements Authenticator {
    private static final Logger logger = Logger.getLogger(SmsOtpAuthenticator.class.getName());

    @Override
    public void authenticate(AuthenticationFlowContext context) {
        UserModel user = context.getUser();
        String phoneNumber = user.getFirstAttribute("phoneNumber");

        if (phoneNumber == null || phoneNumber.isEmpty()) {
            context.failure(AuthenticationFlowError.INTERNAL_ERROR);
            return;
        }

        //6자리 otp 생성
        String otp = String.format("%06d", new Random().nextInt(1000000));

        //otp 를 세션에 저장
        AuthenticationSessionModel session = context.getAuthenticationSession();
        session.setAuthNote("SMS_OTP", otp);
        logger.info("Generated OTP for user " + user.getUsername() + ": " + otp);

        //otp 입력화면으로 이동
        context.challenge(context.form().createForm("sms-otp.ftl"));
    }

    @Override
    public void action(AuthenticationFlowContext context) {
        String enteredOtp = context.getHttpRequest().getDecodedFormParameters().getFirst("otp");
        String expectedOtp = context.getAuthenticationSession().getAuthNote("SMS_OTP");

        logger.info("action...enteredOtp=[" + enteredOtp + "], expectedOtp=[" + expectedOtp + "]");

        if (enteredOtp != null && enteredOtp.equals(expectedOtp)) {
            context.success();
        } else {
//            context.failure(AuthenticationFlowError.INVALID_CREDENTIALS);
            context.challenge(context.form()
                    .setError("Invalid OTP")
                    .createForm("sms-otp.ftl"));
        }
    }

    @Override
    public boolean requiresUser() {
        return false;
    }

    @Override
    public boolean configuredFor(KeycloakSession session, RealmModel realm, UserModel user) {
        return false;
    }

    @Override
    public void setRequiredActions(KeycloakSession session, RealmModel realm, UserModel user) {

    }

    @Override
    public void close() {

    }
}
