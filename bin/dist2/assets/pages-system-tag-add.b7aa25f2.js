import{_ as e,D as a,N as t,O as s,s as n,H as i,y as l,o,c as r,w as d,i as u,a as m,d as f,f as c,g as p,Q as h,q as b,X as g}from"./index-eb11052f.js";import{_}from"./uni-easyinput.616bd617.js";import{_ as y}from"./uni-forms-item.08cc8d8e.js";import{_ as D}from"./uni-forms.aae0e4e3.js";import{v}from"./uni-id-tag.4864df9f.js";const V=a.database();V.command;function x(e){let a={};for(let t in v)e.includes(t)&&(a[t]=v[t]);return a}const j=e({data(){let e={tagid:"",name:"",description:""};return{formData:e,formOptions:{},rules:{...x(Object.keys(e))}}},onReady(){this.$refs.form.setRules(this.rules)},methods:{submit(){t({mask:!0}),this.$refs.form.validate().then((e=>this.submitForm(e))).catch((()=>{})).finally((()=>{s()}))},submitForm(e){return V.collection("uni-id-tag").add(e).then((e=>{n({title:"新增成功"}),this.getOpenerEventChannel().emit("refreshData"),this.getOpenerEventChannel().emit("refreshCheckboxData"),setTimeout((()=>i()),500)})).catch((e=>{l({content:e.message||"请求服务失败",showCancel:!1})}))}}},[["render",function(e,a,t,s,n,i){const l=c(p("uni-easyinput"),_),v=c(p("uni-forms-item"),y),V=h,x=b,j=g,k=u,C=c(p("uni-forms"),D);return o(),r(k,{class:"uni-container"},{default:d((()=>[m(C,{ref:"form",value:n.formData,validateTrigger:"bind"},{default:d((()=>[m(v,{name:"tagid",label:"标签tagid",required:""},{default:d((()=>[m(l,{placeholder:"标签的tagid",modelValue:n.formData.tagid,"onUpdate:modelValue":a[0]||(a[0]=e=>n.formData.tagid=e)},null,8,["modelValue"])])),_:1}),m(v,{name:"name",label:"标签名称",required:""},{default:d((()=>[m(l,{placeholder:"标签名称",modelValue:n.formData.name,"onUpdate:modelValue":a[1]||(a[1]=e=>n.formData.name=e)},null,8,["modelValue"])])),_:1}),m(v,{name:"description",label:"标签描述"},{default:d((()=>[m(V,{placeholder:"标签描述",onInput:a[2]||(a[2]=a=>e.binddata("description",a.detail.value)),class:"uni-textarea-border",modelValue:n.formData.description,"onUpdate:modelValue":a[3]||(a[3]=e=>n.formData.description=e)},null,8,["modelValue"])])),_:1}),m(k,{class:"uni-button-group"},{default:d((()=>[m(x,{type:"primary",class:"uni-button",style:{width:"100px"},onClick:i.submit},{default:d((()=>[f("提交")])),_:1},8,["onClick"]),m(j,{"open-type":"navigateBack",style:{"margin-left":"15px"}},{default:d((()=>[m(x,{class:"uni-button",style:{width:"100px"}},{default:d((()=>[f("返回")])),_:1})])),_:1})])),_:1})])),_:1},8,["value"])])),_:1})}],["__scopeId","data-v-cb3ef994"]]);export{j as default};
