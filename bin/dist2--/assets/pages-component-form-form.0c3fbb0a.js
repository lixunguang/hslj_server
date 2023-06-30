import{_ as e,s as r,o as t,c as a,w as n,i as s,a as u,b as l,r as i,d as o,e as c,I as m,E as f,f as d,C as h,G as g,B as k,D as b,l as _,k as p,v as M,H as R}from"./index-886f4693.js";const w={error:"",check:function(e,r){for(var t=0;t<r.length;t++){if(!r[t].checkType)return!0;if(!r[t].name)return!0;if(!r[t].errorMsg)return!0;if(!e[r[t].name])return this.error=r[t].errorMsg,!1;switch(r[t].checkType){case"string":if(!new RegExp("^.{"+r[t].checkRule+"}$").test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"int":if(!new RegExp("^(-[1-9]|[1-9])[0-9]{"+r[t].checkRule+"}$").test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"between":if(!this.isNumber(e[r[t].name]))return this.error=r[t].errorMsg,!1;if((a=r[t].checkRule.split(","))[0]=Number(a[0]),a[1]=Number(a[1]),e[r[t].name]>a[1]||e[r[t].name]<a[0])return this.error=r[t].errorMsg,!1;break;case"betweenD":if(!/^-?[1-9][0-9]?$/.test(e[r[t].name]))return this.error=r[t].errorMsg,!1;if((a=r[t].checkRule.split(","))[0]=Number(a[0]),a[1]=Number(a[1]),e[r[t].name]>a[1]||e[r[t].name]<a[0])return this.error=r[t].errorMsg,!1;break;case"betweenF":var a;if(!/^-?[0-9][0-9]?.+[0-9]+$/.test(e[r[t].name]))return this.error=r[t].errorMsg,!1;if((a=r[t].checkRule.split(","))[0]=Number(a[0]),a[1]=Number(a[1]),e[r[t].name]>a[1]||e[r[t].name]<a[0])return this.error=r[t].errorMsg,!1;break;case"same":if(e[r[t].name]!=r[t].checkRule)return this.error=r[t].errorMsg,!1;break;case"notsame":if(e[r[t].name]==r[t].checkRule)return this.error=r[t].errorMsg,!1;break;case"email":if(!/^\w+([-+.']\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/.test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"phoneno":if(!/^1[0-9]{10,10}$/.test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"zipcode":if(!/^[0-9]{6}$/.test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"reg":if(!new RegExp(r[t].checkRule).test(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"in":if(-1==r[t].checkRule.indexOf(e[r[t].name]))return this.error=r[t].errorMsg,!1;break;case"notnull":if(null==e[r[t].name]||e[r[t].name].length<1)return this.error=r[t].errorMsg,!1}}return!0},isNumber:function(e){return/^-?[1-9][0-9]?.?[0-9]*$/.test(e)}};const v=e({data:()=>({}),methods:{formSubmit:function(e){console.log("form发生了submit事件，携带数据为："+JSON.stringify(e.detail.value));var t=e.detail.value,a=w.check(t,[{name:"nickname",checkType:"string",checkRule:"1,3",errorMsg:"姓名应为1-3个字符"},{name:"gender",checkType:"in",checkRule:"男,女",errorMsg:"请选择性别"},{name:"loves",checkType:"notnull",checkRule:"",errorMsg:"请选择爱好"}]);r(a?{title:"验证通过!",icon:"none"}:{title:w.error,icon:"none"})},formReset:function(e){console.log("清空数据")}}},[["render",function(e,r,w,v,y,N){const $=i(c("page-head"),o),x=s,S=m,T=f,E=d,D=h,I=g,O=k,j=b,z=_,B=p,C=M,F=R;return t(),a(x,null,{default:n((()=>[u($,{title:"form"}),u(x,{class:"uni-padding-wrap uni-common-mt"},{default:n((()=>[u(F,{onSubmit:N.formSubmit,onReset:N.formReset},{default:n((()=>[u(x,{class:"uni-form-item uni-column"},{default:n((()=>[u(x,{class:"title"},{default:n((()=>[l("姓名")])),_:1}),u(S,{class:"uni-input",name:"nickname",placeholder:"请输入姓名"})])),_:1}),u(x,{class:"uni-form-item uni-column"},{default:n((()=>[u(x,{class:"title"},{default:n((()=>[l("性别")])),_:1}),u(I,{name:"gender"},{default:n((()=>[u(D,null,{default:n((()=>[u(T,{value:"男"}),u(E,null,{default:n((()=>[l("男")])),_:1})])),_:1}),u(D,null,{default:n((()=>[u(T,{value:"女"}),u(E,null,{default:n((()=>[l("女")])),_:1})])),_:1})])),_:1})])),_:1}),u(x,{class:"uni-form-item uni-column"},{default:n((()=>[u(x,{class:"title"},{default:n((()=>[l("爱好")])),_:1}),u(j,{name:"loves"},{default:n((()=>[u(D,null,{default:n((()=>[u(O,{value:"读书"}),u(E,null,{default:n((()=>[l("读书")])),_:1})])),_:1}),u(D,null,{default:n((()=>[u(O,{value:"写字"}),u(E,null,{default:n((()=>[l("写字")])),_:1})])),_:1})])),_:1})])),_:1}),u(x,{class:"uni-form-item uni-column"},{default:n((()=>[u(x,{class:"title"},{default:n((()=>[l("年龄")])),_:1}),u(z,{value:"20",name:"age","show-value":""})])),_:1}),u(x,{class:"uni-form-item uni-column"},{default:n((()=>[u(x,{class:"title"},{default:n((()=>[l("保留选项")])),_:1}),u(x,null,{default:n((()=>[u(B,{name:"switch"})])),_:1})])),_:1}),u(x,{class:"uni-btn-v"},{default:n((()=>[u(C,{"form-type":"submit"},{default:n((()=>[l("Submit")])),_:1}),u(C,{type:"default","form-type":"reset"},{default:n((()=>[l("Reset")])),_:1})])),_:1})])),_:1},8,["onSubmit","onReset"])])),_:1})])),_:1})}],["__scopeId","data-v-1f45e1fa"]]);export{v as default};
