import{_ as e,a7 as t,n as s,f as a,g as n,a2 as r,o as l,c as i,w as o,e as p,F as u,a as c,d as g,r as d,t as m,b as f,U as h,p as A,V as _,W as k,i as x}from"./index-eb11052f.js";import{_ as v}from"./uni-popup-dialog.dbe43988.js";let y=()=>console.log("为定义");const C=e({name:"uni-agreements",computed:{agreements(){if(!t.agreements)return[];let{serviceUrl:e,privacyUrl:s}=t.agreements;return[{url:e,title:"用户服务协议"},{url:s,title:"隐私政策条款"}]}},props:{scope:{type:String,default:()=>"register"}},methods:{popupConfirm(){this.isAgree=!0,y()},popup(e){this.needPopupAgreements=!0,this.$nextTick((()=>{e&&(y=e),this.$refs.popupAgreement.open()}))},navigateTo({url:e,title:t}){s({url:"/uni_modules/uni-id-pages/pages/common/webview/webview?url="+e+"&title="+t,success:e=>{},fail:()=>{},complete:()=>{}})},hasAnd:(e,t)=>e.length-1>t,setAgree(e){this.isAgree=!this.isAgree,this.$emit("setAgree",this.isAgree)}},created(){var e,s;this.needAgreements=((null==(s=null==(e=t)?void 0:e.agreements)?void 0:s.scope)||[]).includes(this.scope)},data:()=>({isAgree:!1,needAgreements:!0,needPopupAgreements:!1})},[["render",function(e,t,s,y,C,b){const w=h,T=A,P=_,U=k,$=x,j=a(n("uni-popup-dialog"),v),F=a(n("uni-popup"),r);return b.agreements.length?(l(),i($,{key:0,class:"root"},{default:o((()=>[C.needAgreements?(l(),p(u,{key:0},[c(U,{onChange:b.setAgree},{default:o((()=>[c(P,{class:"checkbox-box"},{default:o((()=>[c(w,{checked:C.isAgree,style:{transform:"scale(0.5)","margin-right":"-6px"}},null,8,["checked"]),c(T,{class:"text"},{default:o((()=>[g("同意")])),_:1})])),_:1})])),_:1},8,["onChange"]),c($,{class:"content"},{default:o((()=>[(l(!0),p(u,null,d(b.agreements,((e,t)=>(l(),i($,{class:"item",key:t},{default:o((()=>[c(T,{class:"agreement text",onClick:t=>b.navigateTo(e)},{default:o((()=>[g(m(e.title),1)])),_:2},1032,["onClick"]),b.hasAnd(b.agreements,t)?(l(),i(T,{key:0,class:"text and",space:"nbsp"},{default:o((()=>[g(" 和 ")])),_:1})):f("",!0)])),_:2},1024)))),128))])),_:1})],64)):f("",!0),C.needAgreements||C.needPopupAgreements?(l(),i(F,{key:1,ref:"popupAgreement",type:"center"},{default:o((()=>[c(j,{confirmText:"同意",onConfirm:b.popupConfirm},{default:o((()=>[c($,{class:"content"},{default:o((()=>[c(T,{class:"text"},{default:o((()=>[g("请先阅读并同意")])),_:1}),(l(!0),p(u,null,d(b.agreements,((e,t)=>(l(),i($,{class:"item",key:t},{default:o((()=>[c(T,{class:"agreement text",onClick:t=>b.navigateTo(e)},{default:o((()=>[g(m(e.title),1)])),_:2},1032,["onClick"]),b.hasAnd(b.agreements,t)?(l(),i(T,{key:0,class:"text and",space:"nbsp"},{default:o((()=>[g(" 和 ")])),_:1})):f("",!0)])),_:2},1024)))),128))])),_:1})])),_:1},8,["onConfirm"])])),_:1},512)):f("",!0)])),_:1})):f("",!0)}],["__scopeId","data-v-2709ff10"]]);export{C as _};
