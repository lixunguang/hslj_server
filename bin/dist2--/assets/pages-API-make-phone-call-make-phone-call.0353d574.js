import{_ as a,aO as e,o as n,c as t,w as l,i as s,a as i,b as u,r as d,d as o,e as c,I as m,v as p}from"./index-886f4693.js";const r=a({data:()=>({title:"makePhoneCall",disabled:!0}),methods:{bindInput:function(a){this.inputValue=a.detail.value,this.inputValue.length>0?this.disabled=!1:this.disabled=!0},makePhoneCall:function(){e({phoneNumber:this.inputValue,success:()=>{console.log("成功拨打电话")}})}}},[["render",function(a,e,r,h,b,f){const _=d(c("page-head"),o),I=s,k=m,C=p;return n(),t(I,null,{default:l((()=>[i(_,{title:b.title},null,8,["title"]),i(I,{class:"uni-padding-wrap uni-common-mt"},{default:l((()=>[i(I,{class:"uni-hello-text uni-center"},{default:l((()=>[u("请在下方输入电话号码")])),_:1}),i(k,{class:"input uni-common-mt",type:"number",name:"input",onInput:f.bindInput},null,8,["onInput"]),i(I,{class:"uni-btn-v uni-common-mt"},{default:l((()=>[i(C,{onClick:f.makePhoneCall,type:"primary",disabled:b.disabled},{default:l((()=>[u("拨打")])),_:1},8,["onClick","disabled"])])),_:1})])),_:1})])),_:1})}],["__scopeId","data-v-58a8c0c5"]]);export{r as default};
