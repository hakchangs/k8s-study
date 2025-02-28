<#import "template.ftl" as layout>

<@layout.registrationLayout displayInfo=true; section>
    <#if section = "header">
        <title>SMS OTP 인증</title>
    <#elseif section = "form">
        <form action="${url.loginAction}" method="post">
            <div class="form-group">
                <label for="otp">인증 코드</label>
                <input type="text" id="otp" name="otp" class="form-control" autofocus>
            </div>
            <div class="form-group">
                <button type="submit" class="btn btn-primary">인증</button>
            </div>
        </form>
    </#if>
</@layout.registrationLayout>