import{_ as e,ar as t,a3 as a,a8 as n,s,ac as r,o as i,c as l,w as c,i as o,a as d,b as h,t as u,g as p,f,x as m}from"./index-886f4693.js";var b=/^<([-A-Za-z0-9_]+)((?:\s+[a-zA-Z_:][-a-zA-Z0-9_:.]*(?:\s*=\s*(?:(?:"[^"]*")|(?:'[^']*')|[^>\s]+))?)*)\s*(\/?)>/,g=/^<\/([-A-Za-z0-9_]+)[^>]*>/,v=/([a-zA-Z_:][-a-zA-Z0-9_:.]*)(?:\s*=\s*(?:(?:"((?:\\.|[^"])*)")|(?:'((?:\\.|[^'])*)')|([^>\s]+)))?/g,x=k("area,base,basefont,br,col,frame,hr,img,input,link,meta,param,embed,command,keygen,source,track,wbr"),_=k("a,address,article,applet,aside,audio,blockquote,button,canvas,center,dd,del,dir,div,dl,dt,fieldset,figcaption,figure,footer,form,frameset,h1,h2,h3,h4,h5,h6,header,hgroup,hr,iframe,isindex,li,map,menu,noframes,noscript,object,ol,output,p,pre,section,script,table,tbody,td,tfoot,th,thead,tr,ul,video"),y=k("abbr,acronym,applet,b,basefont,bdo,big,br,button,cite,code,del,dfn,em,font,i,iframe,img,input,ins,kbd,label,map,object,q,s,samp,script,select,small,span,strike,strong,sub,sup,textarea,tt,u,var"),w=k("colgroup,dd,dt,li,options,p,td,tfoot,th,thead,tr"),O=k("checked,compact,declare,defer,disabled,ismap,multiple,nohref,noresize,noshade,nowrap,readonly,selected"),S=k("script,style");function k(e){for(var t={},a=e.split(","),n=0;n<a.length;n++)t[a[n]]=!0;return t}function A(e){e=function(e){return e.replace(/<\?xml.*\?>\n/,"").replace(/<!doctype.*>\n/,"").replace(/<!DOCTYPE.*>\n/,"")}(e);var t=[],a={node:"root",children:[]};return function(e,t){var a,n,s,r=[],i=e;for(r.last=function(){return this[this.length-1]};e;){if(n=!0,r.last()&&S[r.last()])e=e.replace(new RegExp("([\\s\\S]*?)</"+r.last()+"[^>]*>"),(function(e,a){return a=a.replace(/<!--([\s\S]*?)-->|<!\[CDATA\[([\s\S]*?)]]>/g,"$1$2"),t.chars&&t.chars(a),""})),o("",r.last());else if(0==e.indexOf("\x3c!--")?(a=e.indexOf("--\x3e"))>=0&&(t.comment&&t.comment(e.substring(4,a)),e=e.substring(a+3),n=!1):0==e.indexOf("</")?(s=e.match(g))&&(e=e.substring(s[0].length),s[0].replace(g,o),n=!1):0==e.indexOf("<")&&(s=e.match(b))&&(e=e.substring(s[0].length),s[0].replace(b,c),n=!1),n){var l=(a=e.indexOf("<"))<0?e:e.substring(0,a);e=a<0?"":e.substring(a),t.chars&&t.chars(l)}if(e==i)throw"Parse Error: "+e;i=e}function c(e,a,n,s){if(a=a.toLowerCase(),_[a])for(;r.last()&&y[r.last()];)o("",r.last());if(w[a]&&r.last()==a&&o("",a),(s=x[a]||!!s)||r.push(a),t.start){var i=[];n.replace(v,(function(e,t){var a=arguments[2]?arguments[2]:arguments[3]?arguments[3]:arguments[4]?arguments[4]:O[t]?t:"";i.push({name:t,value:a,escaped:a.replace(/(^|[^\\])"/g,'$1\\"')})})),t.start&&t.start(a,i,s)}}function o(e,a){if(a)for(n=r.length-1;n>=0&&r[n]!=a;n--);else var n=0;if(n>=0){for(var s=r.length-1;s>=n;s--)t.end&&t.end(r[s]);r.length=n}}o()}(e,{start:function(e,n,s){var r={name:e};if(0!==n.length&&(r.attrs=function(e){return e.reduce((function(e,t){var a=t.value,n=t.name;return e[n]?e[n]=e[n]+" "+a:e[n]=a,e}),{})}(n)),s){var i=t[0]||a;i.children||(i.children=[]),i.children.push(r)}else t.unshift(r)},end:function(e){var n=t.shift();if(n.name!==e&&console.error("invalid state: mismatch end tag"),0===t.length)a.children.push(n);else{var s=t[0];s.children||(s.children=[]),s.children.push(n)}},chars:function(e){var n={type:"text",text:e};if(0===t.length)a.children.push(n);else{var s=t[0];s.children||(s.children=[]),s.children.push(n)}},comment:function(e){var a={node:"comment",text:e},n=t[0];n.children||(n.children=[]),n.children.push(a)}}),a.children}const N="/pages/template/list2detail-detail/list2detail-detail";const z=e({data:()=>({title:"",banner:{},htmlNodes:[]}),onLoad(e){const a=e.detailDate||e.payload;try{this.banner=JSON.parse(decodeURIComponent(a))}catch(n){this.banner=JSON.parse(a)}t({title:this.banner.title}),this.getDetail()},onShareAppMessage(){return{title:this.banner.title,path:N+"?detailDate="+JSON.stringify(this.banner)}},onNavigationBarButtonTap(e){0===e.index&&a({service:"share",success:e=>{if(e.provider&&e.provider.length&&~e.provider.indexOf("weixin")){const t=function(e){let t=[];for(let a=0,n=e.length;a<n;a++)"weixin"===e[a]&&(t.push({text:"分享到微信好友",id:"weixin",sort:0}),t.push({text:"分享到微信朋友圈",id:"weixin",sort:1}));return t.sort((function(e,t){return e.sort-t.sort})),t}(e.provider);n({itemList:t.map((e=>e.text)),success:e=>{const t=e.tapIndex;uni.share({provider:"weixin",type:0,title:this.banner.title,scene:0===t?"WXSceneSession":"WXSenceTimeline",href:"https://uniapp.dcloud.io/h5"+N+"?detailDate="+JSON.stringify(this.banner),imageUrl:"https://vkceyugu.cdn.bspapp.com/VKCEYUGU-dc-site/b6304f00-5168-11eb-bd01-97bc1429a9ff.png"})}})}else s({title:"未检测到可用的微信分享服务"})},fail:e=>{s({title:"获取分享服务失败"})}})},methods:{getDetail(){r({url:"https://unidemo.dcloud.net.cn/api/news/36kr/"+this.banner.post_id,success:e=>{if(200==e.statusCode){var t=e.data.content.replace(/\\/g,"").replace(/<img/g,'<img style="display:none;"');this.htmlNodes=A(t)}},fail:()=>{console.log("fail")}})}}},[["render",function(e,t,a,n,s,r){const b=p,g=o,v=f,x=m;return i(),l(g,null,{default:c((()=>[d(g,{class:"banner"},{default:c((()=>[d(b,{class:"banner-img",src:s.banner.cover},null,8,["src"]),d(g,{class:"banner-title"},{default:c((()=>[h(u(s.banner.title),1)])),_:1})])),_:1}),d(g,{class:"article-meta"},{default:c((()=>[d(v,{class:"article-author"},{default:c((()=>[h(u(s.banner.author_name),1)])),_:1}),d(v,{class:"article-text"},{default:c((()=>[h("发表于")])),_:1}),d(v,{class:"article-time"},{default:c((()=>[h(u(s.banner.published_at),1)])),_:1})])),_:1}),d(g,{class:"article-content"},{default:c((()=>[d(x,{nodes:s.htmlNodes},null,8,["nodes"])])),_:1})])),_:1})}],["__scopeId","data-v-aed21ae3"]]);export{z as default};
