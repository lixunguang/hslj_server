import{_ as t,y as a,x as e,s,o as l,c as d,w as n,i as o,K as r,a as i,d as p,p as u,q as c}from"./index-eb11052f.js";const f=t({data:()=>({id:""}),onLoad({id:t}){this.id=t},methods:{publish(){this.id?this.$request("createPublishHtml",{id:this.id},{functionName:"uni-portal",showModal:!1}).then((t=>{"download"in document.createElement("a")?function(t,a){var e=document.createElement("a");e.download=a,e.style.display="none";var s=new Blob([t]);e.href=URL.createObjectURL(s),document.body.appendChild(e),e.click(),document.body.removeChild(e)}(t.body,"index.html"):s({icon:"error",title:"浏览器不支持",duration:800})})).catch((t=>{a({content:t.errMsg,showCancel:!1})})):a({content:"页面出错，请返回重进",showCancel:!1,success(t){e({url:"/pages/system/app/list"})}})}}},[["render",function(t,a,e,s,f,h){const _=u,x=o,m=c;return l(),d(x,{class:"uni-container"},{default:n((()=>[r("h3",{class:"text-separated",style:{padding:"0 0 20rpx 0"}},"步骤1：了解“统一发布页”"),i(x,{style:{"margin-top":"20rpx"}},{default:n((()=>[i(x,{class:"text-separated"},{default:n((()=>[i(_,{class:"strong"},{default:n((()=>[p("uni-portal ")])),_:1}),i(_,null,{default:n((()=>[p("是 uni-app 提供的一套开箱即用的“统一发布页”。")])),_:1})])),_:1}),i(x,{class:"text-separated"},{default:n((()=>[i(_,{class:"strong"},{default:n((()=>[p("uni-portal ")])),_:1}),i(_,null,{default:n((()=>[p("可作为面向用户的统一业务名片，在一个页面集中展现：App下载地址、小程序二维码、H5访问链接等信息。")])),_:1})])),_:1}),i(x,{class:"text-separated"},{default:n((()=>[i(_,{style:{"font-size":"16px"}},{default:n((()=>[p("uni-app 官方示例的发布页就是基于"),i(_,{class:"strong"},{default:n((()=>[p("uni-portal ")])),_:1}),p(" 制作的，"),r("a",{href:"https://hellouniapp.dcloud.net.cn/portal",target:"_blank",class:"a-label"},"点击体验")])),_:1})])),_:1})])),_:1}),r("h3",{class:"text-separated",style:{padding:"40rpx 0 20rpx 0"}},"步骤2：获取“统一发布页”"),i(x,{class:"flex text-separated",style:{"margin-top":"20rpx"}},{default:n((()=>[i(_,null,{default:n((()=>[i(x,{class:"strong"},{default:n((()=>[p("uni-portal ")])),_:1}),p(" 可根据「应用管理」中所填写的应用信息，一键生成发布页： ")])),_:1}),i(m,{class:"custom-button",size:"mini",type:"primary",onClick:h.publish,style:{margin:"0"}},{default:n((()=>[p("生成并下载发布页")])),_:1},8,["onClick"])])),_:1}),r("h3",{class:"text-separated",style:{padding:"40rpx 0 20rpx 0"}},"步骤3：上传“统一发布页”"),i(x,{style:{"margin-top":"20rpx"}},{default:n((()=>[i(x,{class:"text-separated"},{default:n((()=>[i(_,null,{default:n((()=>[p(" 步骤2下载的“统一发布页”，是一个静态HTML页面，你可以直接在本地浏览器中打开访问。 ")])),_:1})])),_:1}),i(x,{class:"text-separated"},{default:n((()=>[i(_,null,{default:n((()=>[p(" 为了让用户访问到这个“统一发布页”，你需要将该静态HTML文件上传到你的服务器中；推荐使用"),r("a",{href:"https://uniapp.dcloud.io/uniCloud/hosting",target:"_blank",class:"a-label",style:{padding:"5px"}},"前端网页托管"),p("，因为前端网页托管具备使用更简单、价格更便宜、访问更快等优点。 ")])),_:1})])),_:1})])),_:1})])),_:1})}],["__scopeId","data-v-9817eb0f"]]);export{f as default};
