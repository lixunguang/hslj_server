import{_ as a,W as e,at as t,ah as n,au as l,av as i,aw as o,ax as s,s as c,o as u,c as r,w as d,i as g,a as p,b as f,m,r as k,d as v,e as w,v as C}from"./index-886f4693.js";const h="/pages/extUI/calendar/calendar";const _=a({data:()=>({title:"navigate"}),computed:{...e({hasLeftWin:a=>!a.noMatchLeftWindow})},methods:{navigateTo(){t({url:"new-page/new-vue-page-1?data=Hello"})},navigateBack(){n()},redirectTo(){l({url:"new-page/new-vue-page-1"})},switchTab(){i({url:"/pages/tabBar/template/template"})},reLaunch(){this.hasLeftWin?o({url:"/pages/component/view/view"}):o({url:"/pages/tabBar/component/component"})},customAnimation(){t({url:"new-page/new-vue-page-1?data=使用自定义动画打开页面",animationType:"slide-in-bottom",animationDuration:200})},preloadPage(){s({url:h,success(){c({title:"页面预载成功"})},fail(){c({title:"页面预载失败"})}})},unPreloadPage(){uni.unPreloadPage({url:h})},navigateToPreloadPage(){t({url:h})}}},[["render",function(a,e,t,n,l,i){const o=k(w("page-head"),v),s=C,c=g;return u(),r(c,null,{default:d((()=>[p(o,{title:l.title},null,8,["title"]),p(c,{class:"uni-padding-wrap uni-common-mt"},{default:d((()=>[p(c,{class:"uni-btn-v"},{default:d((()=>[p(s,{onClick:i.navigateTo},{default:d((()=>[f("跳转新页面，并传递数据")])),_:1},8,["onClick"]),p(s,{onClick:i.navigateBack},{default:d((()=>[f("返回上一页")])),_:1},8,["onClick"]),p(s,{onClick:i.redirectTo},{default:d((()=>[f("在当前页面打开")])),_:1},8,["onClick"]),p(s,{onClick:i.switchTab},{default:d((()=>[f("切换到模板选项卡")])),_:1},8,["onClick"]),a.hasLeftWin?m("",!0):(u(),r(s,{key:0,onClick:i.reLaunch},{default:d((()=>[f("关闭所有页面，打开首页")])),_:1},8,["onClick"])),p(s,{onClick:i.preloadPage},{default:d((()=>[f("预载复杂页面")])),_:1},8,["onClick"]),p(s,{onClick:i.navigateToPreloadPage},{default:d((()=>[f("打开复杂页面")])),_:1},8,["onClick"])])),_:1})])),_:1})])),_:1})}]]);export{_ as default};
