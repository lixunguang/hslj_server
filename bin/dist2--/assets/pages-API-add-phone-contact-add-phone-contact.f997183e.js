import{_ as a,aL as e,Z as l,o as n,c as t,w as s,i as u,a as i,b as c,r as o,d,e as f,I as p,v as h}from"./index-886f4693.js";const m=a({data:()=>({title:"addPhoneContact",name:"",phone:""}),methods:{nameChange:function(a){this.name=a.detail.value},phoneChange:function(a){this.phone=a.detail.value},async add(){e({firstName:this.name,mobilePhoneNumber:this.phone,success:function(){l({content:"已成功添加联系人！",showCancel:!1})},fail:function(){l({content:"添加联系人失败！",showCancel:!1})}})}}},[["render",function(a,e,l,m,r,_){const b=o(f("page-head"),d),C=u,v=p,g=h;return n(),t(C,null,{default:s((()=>[i(b,{title:r.title},null,8,["title"]),i(C,{class:"uni-common-mt"},{default:s((()=>[i(C,{class:"uni-list"},{default:s((()=>[i(C,{class:"uni-list-cell"},{default:s((()=>[i(C,{class:"uni-list-cell-left"},{default:s((()=>[i(C,{class:"uni-label"},{default:s((()=>[c("名称")])),_:1})])),_:1}),i(C,{class:"uni-list-cell-db"},{default:s((()=>[i(v,{class:"uni-input",type:"text",placeholder:"请输入联系人名称",name:"name",value:r.name,onInput:_.nameChange},null,8,["value","onInput"])])),_:1})])),_:1}),i(C,{class:"uni-list-cell"},{default:s((()=>[i(C,{class:"uni-list-cell-left"},{default:s((()=>[i(C,{class:"uni-label"},{default:s((()=>[c("手机号")])),_:1})])),_:1}),i(C,{class:"uni-list-cell-db"},{default:s((()=>[i(v,{class:"uni-input",type:"text",placeholder:"请输入联系人手机号",name:"phone",value:r.phone,onInput:_.phoneChange},null,8,["value","onInput"])])),_:1})])),_:1})])),_:1}),i(C,{class:"uni-padding-wrap"},{default:s((()=>[i(C,{class:"uni-btn-v"},{default:s((()=>[i(g,{type:"primary",class:"btn-setstorage",onClick:_.add},{default:s((()=>[c("添加联系人")])),_:1},8,["onClick"])])),_:1})])),_:1})])),_:1})])),_:1})}]]);export{m as default};
