import{_ as l,s as o,o as e,c as s,w as c,i as t,a as n,b as i,r,d as a,e as u,f as d,S as m}from"./index-886f4693.js";const f=l({data:()=>({scrollTop:0,old:{scrollTop:0}}),methods:{upper:function(l){console.log(l)},lower:function(l){console.log(l)},scroll:function(l){console.log(l),this.old.scrollTop=l.detail.scrollTop},goTop:function(l){this.scrollTop=this.old.scrollTop,this.$nextTick((function(){this.scrollTop=0})),o({icon:"none",title:"纵向滚动 scrollTop 值已被修改为 0"})}}},[["render",function(l,o,f,p,_,g){const w=r(u("page-head"),a),T=d,S=t,b=m;return e(),s(S,null,{default:c((()=>[n(w,{title:"scroll-view,区域滚动视图"}),n(S,{class:"uni-padding-wrap uni-common-mt"},{default:c((()=>[n(S,{class:"uni-title uni-common-mt"},{default:c((()=>[i(" Vertical Scroll "),n(T,null,{default:c((()=>[i("\\n纵向滚动")])),_:1})])),_:1}),n(S,null,{default:c((()=>[n(b,{"scroll-top":_.scrollTop,"scroll-y":"true",class:"scroll-Y",onScrolltoupper:g.upper,onScrolltolower:g.lower,onScroll:g.scroll},{default:c((()=>[n(S,{id:"demo1",class:"scroll-view-item uni-bg-red"},{default:c((()=>[i("A")])),_:1}),n(S,{id:"demo2",class:"scroll-view-item uni-bg-green"},{default:c((()=>[i("B")])),_:1}),n(S,{id:"demo3",class:"scroll-view-item uni-bg-blue"},{default:c((()=>[i("C")])),_:1})])),_:1},8,["scroll-top","onScrolltoupper","onScrolltolower","onScroll"])])),_:1}),n(S,{onClick:g.goTop,class:"uni-link uni-center uni-common-mt"},{default:c((()=>[i(" 点击这里返回顶部 ")])),_:1},8,["onClick"]),n(S,{class:"uni-title uni-common-mt"},{default:c((()=>[i(" Horizontal Scroll "),n(T,null,{default:c((()=>[i("\\n横向滚动")])),_:1})])),_:1}),n(S,null,{default:c((()=>[n(b,{class:"scroll-view_H","scroll-x":"true",onScroll:g.scroll,"scroll-left":"120"},{default:c((()=>[n(S,{id:"demo1",class:"scroll-view-item_H uni-bg-red"},{default:c((()=>[i("A")])),_:1}),n(S,{id:"demo2",class:"scroll-view-item_H uni-bg-green"},{default:c((()=>[i("B")])),_:1}),n(S,{id:"demo3",class:"scroll-view-item_H uni-bg-blue"},{default:c((()=>[i("C")])),_:1})])),_:1},8,["onScroll"])])),_:1}),n(S,{class:"uni-common-pb"})])),_:1})])),_:1})}],["__scopeId","data-v-9191317f"]]);export{f as default};
