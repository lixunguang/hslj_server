import{_ as e,D as a,s as t,n,o as r,c as i,w as s,i as l,a as o,d,t as u,e as c,r as p,F as m,f as g,g as h,M as f,q as _,h as b,j as y,k as C,l as x}from"./index-eb11052f.js";import{_ as w}from"./uni-stat-breadcrumb.d2425bd8.js";import{_ as z}from"./uni-dateformat.6fbc712f.js";import{_ as k}from"./uni-pagination.4e528405.js";import{_ as D}from"./unicloud-db.7bb513e6.js";const j=a.database(),V=["user_id.username","user_id.nickname","type","ip"];const $=e({data:()=>({collectionList:[j.collection("uni-id-log").field("type, ip, create_date, user_id").getTemp(),j.collection("uni-id-users").field("_id, username,nickname").getTemp()],query:"",where:"",orderby:"create_date desc",options:{pageSize:20,pageCurrent:1}}),methods:{getWhere(){const e=this.query.trim();if(!e)return"";let a;try{a=new RegExp(e,"i")}catch(n){return void t({title:"请勿输入等不满足正则格式的符号",icon:"none"})}return V.map((e=>a+".test("+e+")")).join(" || ")},search(){const e=this.getWhere(),a=e===this.where;this.where=e,a&&this.loadData()},loadData(e=!0){this.$refs.udb.loadData({clear:e})},onPageChanged(e){this.$refs.udb.loadData({current:e.current})},navigateTo(e){n({url:e,events:{refreshData:()=>{this.loadData()}}})},selectedItems(){var e=this.$refs.udb.dataList;return this.selectedIndexs.map((a=>e[a]._id))},delTable(){this.$refs.udb.remove(this.selectedItems())},selectionChange(e){this.selectedIndexs=e.detail.index},confirmDelete(e){this.$refs.udb.remove(e)}}},[["render",function(e,a,t,n,j,V){const $=g(h("uni-stat-breadcrumb"),w),q=f,v=_,T=l,I=g(h("uni-th"),b),L=g(h("uni-tr"),y),P=g(h("uni-td"),C),U=g(h("uni-dateformat"),z),S=g(h("uni-table"),x),W=g(h("uni-pagination"),k),E=g(h("unicloud-db"),D);return r(),i(T,null,{default:s((()=>[o(T,{class:"uni-header"},{default:s((()=>[o($,{class:"uni-stat-breadcrumb-on-phone"}),o(T,{class:"uni-group"},{default:s((()=>[o(q,{class:"uni-search",type:"text",modelValue:j.query,"onUpdate:modelValue":a[0]||(a[0]=e=>j.query=e),onConfirm:V.search,placeholder:e.$t("common.placeholder.query")},null,8,["modelValue","onConfirm","placeholder"]),o(v,{class:"uni-button hide-on-phone",type:"default",size:"mini",onClick:V.search},{default:s((()=>[d(u(e.$t("common.button.search")),1)])),_:1},8,["onClick"])])),_:1})])),_:1}),o(T,{class:"uni-container"},{default:s((()=>[o(E,{ref:"udb",collection:j.collectionList,options:j.options,where:j.where,"page-data":"replace",orderby:j.orderby,getcount:!0,"page-size":j.options.pageSize,"page-current":j.options.pageCurrent},{default:s((({data:e,pagination:a,loading:t,error:n})=>[o(S,{loading:t,emptyText:n.message||"没有更多数据",border:"",stripe:""},{default:s((()=>[o(L,null,{default:s((()=>[o(I,{align:"center"},{default:s((()=>[d("序号")])),_:1}),o(I,{align:"center"},{default:s((()=>[d("用户名")])),_:1}),o(I,{align:"center"},{default:s((()=>[d("昵称")])),_:1}),o(I,{align:"center"},{default:s((()=>[d("内容")])),_:1}),o(I,{align:"center"},{default:s((()=>[d("IP")])),_:1}),o(I,{align:"center"},{default:s((()=>[d("时间")])),_:1})])),_:1}),(r(!0),c(m,null,p(e,((e,t)=>(r(),i(L,{key:t},{default:s((()=>[o(P,{align:"center"},{default:s((()=>[d(u((a.current-1)*a.size+(t+1)),1)])),_:2},1024),o(P,{align:"center"},{default:s((()=>[d(u(e.user_id[0]&&e.user_id[0].username||"-"),1)])),_:2},1024),o(P,{align:"center"},{default:s((()=>[d(u(e.user_id[0]&&e.user_id[0].nickname||"-"),1)])),_:2},1024),o(P,{align:"center"},{default:s((()=>[d(u(e.type),1)])),_:2},1024),o(P,{align:"center"},{default:s((()=>[d(u(e.ip),1)])),_:2},1024),o(P,{align:"center"},{default:s((()=>[o(U,{date:e.create_date,threshold:[0,0]},null,8,["date"])])),_:2},1024)])),_:2},1024)))),128))])),_:2},1032,["loading","emptyText"]),o(T,{class:"uni-pagination-box"},{default:s((()=>[o(W,{"show-icon":"","page-size":a.size,modelValue:a.current,"onUpdate:modelValue":e=>a.current=e,total:a.count,onChange:V.onPageChanged},null,8,["page-size","modelValue","onUpdate:modelValue","total","onChange"])])),_:2},1024)])),_:1},8,["collection","options","where","orderby","page-size","page-current"])])),_:1})])),_:1})}]]);export{$ as default};
