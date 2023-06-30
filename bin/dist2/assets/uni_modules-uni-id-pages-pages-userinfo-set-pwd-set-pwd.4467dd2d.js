import{_ as s,D as a,Z as o,a7 as e,s as t,y as r,$ as i,o as n,c as l,w as d,i as u,a as c,d as m,b as p,m as f,p as w,f as h,g,q as P}from"./index-eb11052f.js";import{_}from"./uni-match-media.acbeb623.js";import{_ as y}from"./uni-easyinput.616bd617.js";import{_ as b}from"./uni-forms-item.08cc8d8e.js";import{_ as D}from"./uni-id-pages-sms-form.3b16befc.js";import{_ as k}from"./uni-forms.aae0e4e3.js";import{_ as I}from"./uni-popup-captcha.cd4d537f.js";import{p as C}from"./password.6cc85f39.js";import"./uni-captcha.53f0e573.js";const U=a.importObject("uni-id-co",{customUI:!0});const V=s({name:"set-pwd.vue",data:()=>({uniIdRedirectUrl:"",loginType:"",logo:"/static/logo.png",focusNewPassword:!1,focusNewPassword2:!1,allowSkip:!1,formData:{code:"",captcha:"",newPassword:"",newPassword2:""},rules:C.getPwdRules("newPassword","newPassword2")}),computed:{userInfo:()=>o.userInfo},onReady(){this.$refs.form.setRules(this.rules)},onLoad(s){var a;this.uniIdRedirectUrl=s.uniIdRedirectUrl,this.loginType=s.loginType,e.setPasswordAfterLogin&&(null==(a=e.setPasswordAfterLogin)?void 0:a.allowSkip)&&(this.allowSkip=!0)},methods:{submit(){if(!/^\d{6}$/.test(this.formData.code))return this.$refs.smsCode.focusSmsCodeInput=!0,t({title:"验证码格式不正确",icon:"none"});this.$refs.form.validate().then((s=>{U.setPwd({password:this.formData.newPassword,code:this.formData.code,captcha:this.formData.captcha}).then((s=>{r({content:"密码设置成功",showCancel:!1,success:()=>{i.loginBack({uniIdRedirectUrl:this.uniIdRedirectUrl,loginType:this.loginType})}})})).catch((s=>{r({content:s.message,showCancel:!1})}))})).catch((s=>{"uni-id-captcha-required"==s.errCode?this.$refs.popup.open():console.log(s.errMsg)})).finally((s=>{this.formData.captcha=""}))},skip(){i.loginBack({uniIdRedirectUrl:this.uniIdRedirectUrl})}}},[["render",function(s,a,o,e,t,r){const i=f,C=u,U=w,V=h(g("uni-match-media"),_),j=h(g("uni-easyinput"),y),R=h(g("uni-forms-item"),b),v=h(g("uni-id-pages-sms-form"),D),x=P,B=h(g("uni-forms"),k),N=h(g("uni-popup-captcha"),I);return n(),l(C,{class:"uni-content"},{default:d((()=>[c(V,{"min-width":690},{default:d((()=>[c(C,{class:"login-logo"},{default:d((()=>[c(i,{src:t.logo},null,8,["src"])])),_:1}),c(U,{class:"title title-box"},{default:d((()=>[m("设置密码")])),_:1})])),_:1}),c(B,{class:"set-password-form",ref:"form",value:t.formData,"err-show-type":"toast"},{default:d((()=>[c(U,{class:"tip"},{default:d((()=>[m("输入密码")])),_:1}),c(R,{name:"newPassword"},{default:d((()=>[c(j,{focus:t.focusNewPassword,onBlur:a[0]||(a[0]=s=>t.focusNewPassword=!1),class:"input-box",type:"password",inputBorder:!1,modelValue:t.formData.newPassword,"onUpdate:modelValue":a[1]||(a[1]=s=>t.formData.newPassword=s),placeholder:"请输入密码"},null,8,["focus","modelValue"])])),_:1}),c(U,{class:"tip"},{default:d((()=>[m("再次输入密码")])),_:1}),c(R,{name:"newPassword2"},{default:d((()=>[c(j,{focus:t.focusNewPassword2,onBlur:a[2]||(a[2]=s=>t.focusNewPassword2=!1),class:"input-box",type:"password",inputBorder:!1,modelValue:t.formData.newPassword2,"onUpdate:modelValue":a[3]||(a[3]=s=>t.formData.newPassword2=s),placeholder:"请再次输入新密码"},null,8,["focus","modelValue"])])),_:1}),c(v,{modelValue:t.formData.code,"onUpdate:modelValue":a[4]||(a[4]=s=>t.formData.code=s),type:"set-pwd-by-sms",ref:"smsCode",phone:r.userInfo.mobile},null,8,["modelValue","phone"]),c(C,{class:"link-box"},{default:d((()=>[c(x,{class:"uni-btn send-btn",type:"primary",onClick:r.submit},{default:d((()=>[m("确认")])),_:1},8,["onClick"]),t.allowSkip?(n(),l(x,{key:0,class:"uni-btn send-btn",type:"default",onClick:r.skip},{default:d((()=>[m("跳过")])),_:1},8,["onClick"])):p("",!0)])),_:1})])),_:1},8,["value"]),c(N,{onConfirm:r.submit,modelValue:t.formData.captcha,"onUpdate:modelValue":a[5]||(a[5]=s=>t.formData.captcha=s),scene:"set-pwd-by-sms",ref:"popup"},null,8,["onConfirm","modelValue"])])),_:1})}],["__scopeId","data-v-f34b7681"]]);export{V as default};
