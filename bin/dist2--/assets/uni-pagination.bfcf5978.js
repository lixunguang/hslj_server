import{bd as e,_ as t,r as n,e as i,o as a,c as r,w as s,a as u,b as c,t as p,m as o,L as l,z as g,F as h,A as d,f as x,v as m,N as _,i as f}from"./index-886f4693.js";import{_ as P}from"./uni-icons.072c6819.js";const I={en:{"uni-pagination.prevText":"prev","uni-pagination.nextText":"next","uni-pagination.piecePerPage":"piece/page"},es:{"uni-pagination.prevText":"anterior","uni-pagination.nextText":"prxima","uni-pagination.piecePerPage":"Art��culo/P��gina"},fr:{"uni-pagination.prevText":"précédente","uni-pagination.nextText":"suivante","uni-pagination.piecePerPage":"Articles/Pages"},"zh-Hans":{"uni-pagination.prevText":"上一页","uni-pagination.nextText":"下一页","uni-pagination.piecePerPage":"条/页"},"zh-Hant":{"uni-pagination.prevText":"上一頁","uni-pagination.nextText":"下一頁","uni-pagination.piecePerPage":"條/頁"}},{t:k}=e(I);const v=t({name:"UniPagination",emits:["update:modelValue","input","change","pageSizeChange"],props:{value:{type:[Number,String],default:1},modelValue:{type:[Number,String],default:1},prevText:{type:String},nextText:{type:String},piecePerPageText:{type:String},current:{type:[Number,String],default:1},total:{type:[Number,String],default:0},pageSize:{type:[Number,String],default:10},showIcon:{type:[Boolean,String],default:!1},showPageSize:{type:[Boolean,String],default:!1},pagerCount:{type:Number,default:7},pageSizeRange:{type:Array,default:()=>[20,50,100,500]}},data:()=>({pageSizeIndex:0,currentIndex:1,paperData:[],pickerShow:!1}),computed:{piecePerPage(){return this.piecePerPageText||k("uni-pagination.piecePerPage")},prevPageText(){return this.prevText||k("uni-pagination.prevText")},nextPageText(){return this.nextText||k("uni-pagination.nextText")},maxPage(){let e=1,t=Number(this.total),n=Number(this.pageSize);return t&&n&&(e=Math.ceil(t/n)),e},paper(){const e=this.currentIndex,t=this.pagerCount,n=this.total,i=this.pageSize;let a=[],r=[],s=Math.ceil(n/i);for(let c=0;c<s;c++)a.push(c+1);r.push(1);const u=a[a.length-(t+1)/2];return a.forEach(((n,i)=>{(t+1)/2>=e?n<t+1&&n>1&&r.push(n):e+2<=u?n>e-(t+1)/2&&n<e+(t+1)/2&&r.push(n):(n>e-(t+1)/2||s-t<n)&&n<a[a.length-1]&&r.push(n)})),s>t?((t+1)/2>=e?r[r.length-1]="...":e+2<=u?(r[1]="...",r[r.length-1]="..."):r[1]="...",r.push(a[a.length-1])):(t+1)/2>=e||e+2<=u||(r.shift(),r.push(a[a.length-1])),r}},watch:{current:{immediate:!0,handler(e,t){this.currentIndex=e<1?1:e}},value:{immediate:!0,handler(e){1===Number(this.current)&&(this.currentIndex=e<1?1:e)}},pageSizeIndex(e){this.$emit("pageSizeChange",this.pageSizeRange[e])}},methods:{pickerChange(e){this.pageSizeIndex=e.detail.value,this.pickerClick()},pickerClick(){const e=document.querySelector("body");if(!e)return;const t="uni-pagination-picker-show";this.pickerShow=!this.pickerShow,this.pickerShow?e.classList.add(t):setTimeout((()=>e.classList.remove(t)),300)},selectPage(e,t){if(parseInt(e))this.currentIndex=e,this.change("current");else{let e=Math.ceil(this.total/this.pageSize);if(t<=1)return void(this.currentIndex-5>1?this.currentIndex-=5:this.currentIndex=1);if(t>=6)return void(this.currentIndex+5>e?this.currentIndex=e:this.currentIndex+=5)}},clickLeft(){1!==Number(this.currentIndex)&&(this.currentIndex-=1,this.change("prev"))},clickRight(){Number(this.currentIndex)>=this.maxPage||(this.currentIndex+=1,this.change("next"))},change(e){this.$emit("input",this.currentIndex),this.$emit("update:modelValue",this.currentIndex),this.$emit("change",{type:e,current:this.currentIndex})}}},[["render",function(e,t,I,k,v,S){const y=x,z=n(i("uni-icons"),P),T=m,b=_,C=f;return a(),r(C,{class:"uni-pagination"},{default:s((()=>[!0===I.showPageSize||"true"===I.showPageSize?(a(),r(b,{key:0,class:"select-picker",mode:"selector",value:v.pageSizeIndex,range:I.pageSizeRange,onChange:S.pickerChange,onCancel:S.pickerClick,onClick:S.pickerClick},{default:s((()=>[u(T,{type:"default",size:"mini",plain:!0},{default:s((()=>[u(y,null,{default:s((()=>[c(p(I.pageSizeRange[v.pageSizeIndex])+" "+p(S.piecePerPage),1)])),_:1}),u(z,{class:"select-picker-icon",type:"arrowdown",size:"12",color:"#999"})])),_:1})])),_:1},8,["value","range","onChange","onCancel","onClick"])):o("",!0),u(C,{class:"uni-pagination__total is-phone-hide"},{default:s((()=>[c("共 "+p(I.total)+" 条",1)])),_:1}),u(C,{class:l(["uni-pagination__btn",1===v.currentIndex?"uni-pagination--disabled":"uni-pagination--enabled"]),"hover-class":1===v.currentIndex?"":"uni-pagination--hover","hover-start-time":20,"hover-stay-time":70,onClick:S.clickLeft},{default:s((()=>[!0===I.showIcon||"true"===I.showIcon?(a(),r(z,{key:0,color:"#666",size:"16",type:"left"})):(a(),r(y,{key:1,class:"uni-pagination__child-btn"},{default:s((()=>[c(p(S.prevPageText),1)])),_:1}))])),_:1},8,["class","hover-class","onClick"]),u(C,{class:"uni-pagination__num uni-pagination__num-flex-none"},{default:s((()=>[u(C,{class:"uni-pagination__num-current"},{default:s((()=>[u(y,{class:"uni-pagination__num-current-text is-pc-hide",style:{color:"#409EFF"}},{default:s((()=>[c(p(v.currentIndex),1)])),_:1}),u(y,{class:"uni-pagination__num-current-text is-pc-hide"},{default:s((()=>[c("/"+p(S.maxPage||0),1)])),_:1}),(a(!0),g(h,null,d(S.paper,((e,t)=>(a(),r(C,{key:t,class:l([{"page--active":e===v.currentIndex},"uni-pagination__num-tag tag--active is-phone-hide"]),onClick:n=>S.selectPage(e,t)},{default:s((()=>[u(y,null,{default:s((()=>[c(p(e),1)])),_:2},1024)])),_:2},1032,["class","onClick"])))),128))])),_:1})])),_:1}),u(C,{class:l(["uni-pagination__btn",v.currentIndex>=S.maxPage?"uni-pagination--disabled":"uni-pagination--enabled"]),"hover-class":v.currentIndex===S.maxPage?"":"uni-pagination--hover","hover-start-time":20,"hover-stay-time":70,onClick:S.clickRight},{default:s((()=>[!0===I.showIcon||"true"===I.showIcon?(a(),r(z,{key:0,color:"#666",size:"16",type:"right"})):(a(),r(y,{key:1,class:"uni-pagination__child-btn"},{default:s((()=>[c(p(S.nextPageText),1)])),_:1}))])),_:1},8,["class","hover-class","onClick"])])),_:1})}],["__scopeId","data-v-1c1d7556"]]);export{v as _};
