import{_ as t,Y as a,Z as e,o as n,c as l,w as o,i as s,a as d,b as u,m as i,r as c,d as r,e as m,$ as f,I as p,v as E}from"./index-886f4693.js";const h=t({data:()=>({title:"video",src:"",danmuList:[{text:"第 1s 出现的弹幕",color:"#ff0000",time:1},{text:"第 3s 出现的弹幕",color:"#ff00ff",time:3}],danmuValue:"",showVideo:!1}),onReady:function(t){this.videoContext=a("myVideo"),this.showVideo=!0},methods:{sendDanmu:function(){this.videoContext.sendDanmu({text:this.danmuValue,color:this.getRandomColor()}),this.danmuValue=""},videoErrorCallback:function(t){e({content:t.target.errMsg,showCancel:!1})},getRandomColor:function(){const t=[];for(let a=0;a<3;++a){let a=Math.floor(256*Math.random()).toString(16);a=1==a.length?"0"+a:a,t.push(a)}return"#"+t.join("")}}},[["render",function(t,a,e,h,C,b){const V=c(m("page-head"),r),g=f,_=s,B=p,v=E;return n(),l(_,null,{default:o((()=>[d(V,{title:C.title},null,8,["title"]),C.showVideo?(n(),l(_,{key:0,class:"uni-padding-wrap uni-common-mt"},{default:o((()=>[d(_,null,{default:o((()=>[d(g,{id:"myVideo",src:"https://img.cdn.aliyun.dcloud.net.cn/guide/uniapp/%E7%AC%AC1%E8%AE%B2%EF%BC%88uni-app%E4%BA%A7%E5%93%81%E4%BB%8B%E7%BB%8D%EF%BC%89-%20DCloud%E5%AE%98%E6%96%B9%E8%A7%86%E9%A2%91%E6%95%99%E7%A8%8B@20181126-lite.m4v",onError:b.videoErrorCallback,"danmu-list":C.danmuList,"enable-danmu":"","danmu-btn":"",controls:"",poster:"https://vkceyugu.cdn.bspapp.com/VKCEYUGU-dc-site/b1476d40-4e5f-11eb-b997-9918a5dda011.png"},null,8,["onError","danmu-list"])])),_:1}),d(_,{class:"uni-list uni-common-mt"},{default:o((()=>[d(_,{class:"uni-list-cell"},{default:o((()=>[d(_,null,{default:o((()=>[d(_,{class:"uni-label"},{default:o((()=>[u("弹幕内容")])),_:1})])),_:1}),d(_,{class:"uni-list-cell-db"},{default:o((()=>[d(B,{modelValue:C.danmuValue,"onUpdate:modelValue":a[0]||(a[0]=t=>C.danmuValue=t),class:"uni-input",type:"text",placeholder:"在此处输入弹幕内容"},null,8,["modelValue"])])),_:1})])),_:1})])),_:1}),d(_,{class:"uni-btn-v"},{default:o((()=>[d(v,{onClick:b.sendDanmu,class:"page-body-button"},{default:o((()=>[u("发送弹幕")])),_:1},8,["onClick"])])),_:1})])),_:1})):i("",!0)])),_:1})}],["__scopeId","data-v-7c6cd374"]]);export{h as default};
