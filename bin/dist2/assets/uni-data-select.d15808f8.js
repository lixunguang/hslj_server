import{D as e,af as t,v as a,_ as l,f as s,A as i,g as c,o,c as n,w as r,e as h,t as u,b as d,a as m,z as p,d as f,E as y,F as _,r as b,i as g,p as D,S as C}from"./index-eb11052f.js";const x=l({name:"uni-data-select",mixins:[e.mixinDatacom||{}],props:{localdata:{type:Array,default:()=>[]},value:{type:[String,Number],default:""},modelValue:{type:[String,Number],default:""},label:{type:String,default:""},placeholder:{type:String,default:"请选择"},emptyTips:{type:String,default:"无选项"},clear:{type:Boolean,default:!0},defItem:{type:Number,default:0},disabled:{type:Boolean,default:!1},format:{type:String,default:""}},data:()=>({showSelector:!1,current:"",mixinDatacomResData:[],apps:[],channels:[],cacheKey:"uni-data-select-lastSelectedValue"}),created(){this.debounceGet=this.debounce((()=>{this.query()}),300),this.collection&&!this.localdata.length&&this.debounceGet()},computed:{typePlaceholder(){const e=this.placeholder,t={"opendb-stat-app-versions":"版本","opendb-app-channels":"渠道","opendb-app-list":"应用"}[this.collection];return t?e+t:e},valueCom(){return this.modelValue}},watch:{localdata:{immediate:!0,handler(e,t){Array.isArray(e)&&t!==e&&(this.mixinDatacomResData=e)}},valueCom(e,t){this.initDefVal()},mixinDatacomResData:{immediate:!0,handler(e){e.length&&this.initDefVal()}}},methods:{debounce(e,t=100){let a=null;return function(...l){a&&clearTimeout(a),a=setTimeout((()=>{e.apply(this,l)}),t)}},query(){this.mixinDatacomEasyGet()},onMixinDatacomPropsChange(){this.collection&&this.debounceGet()},initDefVal(){let e="";if(!this.valueCom&&0!==this.valueCom||this.isDisabled(this.valueCom)){let t;if(this.collection&&(t=this.getCache()),t||0===t)e=t;else{let t="";this.defItem>0&&this.defItem<=this.mixinDatacomResData.length&&(t=this.mixinDatacomResData[this.defItem-1].value),e=t}(e||0===e)&&this.emit(e)}else e=this.valueCom;const t=this.mixinDatacomResData.find((t=>t.value===e));this.current=t?this.formatItemName(t):""},isDisabled(e){let t=!1;return this.mixinDatacomResData.forEach((a=>{a.value===e&&(t=a.disable)})),t},clearVal(){this.emit(""),this.collection&&this.removeCache()},change(e){e.disable||(this.showSelector=!1,this.current=this.formatItemName(e),this.emit(e.value))},emit(e){this.$emit("input",e),this.$emit("update:modelValue",e),this.$emit("change",e),this.collection&&this.setCache(e)},toggleSelector(){this.disabled||(this.showSelector=!this.showSelector)},formatItemName(e){let{text:t,value:a,channel_code:l}=e;if(l=l?`(${l})`:"",this.format){let t="";t=this.format;for(let a in e)t=t.replace(new RegExp(`{${a}}`,"g"),e[a]);return t}return this.collection.indexOf("app-list")>0?`${t}(${a})`:t||`未命名${l}`},getLoadData(){return this.mixinDatacomResData},getCurrentCacheKey(){return this.collection},getCache(e=this.getCurrentCacheKey()){return(t(this.cacheKey)||{})[e]},setCache(e,l=this.getCurrentCacheKey()){let s=t(this.cacheKey)||{};s[l]=e,a(this.cacheKey,s)},removeCache(e=this.getCurrentCacheKey()){let l=t(this.cacheKey)||{};delete l[e],a(this.cacheKey,l)}}},[["render",function(e,t,a,l,x,k){const v=g,S=s(c("uni-icons"),i),w=D,R=C;return o(),n(v,{class:"uni-stat__select"},{default:r((()=>[a.label?(o(),h("span",{key:0,class:"uni-label-text hide-on-phone"},u(a.label+"："),1)):d("",!0),m(v,{class:p(["uni-stat-box",{"uni-stat__actived":x.current}])},{default:r((()=>[m(v,{class:p(["uni-select",{"uni-select--disabled":a.disabled}])},{default:r((()=>[m(v,{class:"uni-select__input-box",onClick:k.toggleSelector},{default:r((()=>[x.current?(o(),n(v,{key:0,class:"uni-select__input-text"},{default:r((()=>[f(u(x.current),1)])),_:1})):(o(),n(v,{key:1,class:"uni-select__input-text uni-select__input-placeholder"},{default:r((()=>[f(u(k.typePlaceholder),1)])),_:1})),x.current&&a.clear&&!a.disabled?(o(),n(v,{key:2,onClick:y(k.clearVal,["stop"])},{default:r((()=>[m(S,{type:"clear",color:"#c0c4cc",size:"24"})])),_:1},8,["onClick"])):(o(),n(v,{key:3},{default:r((()=>[m(S,{type:x.showSelector?"top":"bottom",size:"14",color:"#999"},null,8,["type"])])),_:1}))])),_:1},8,["onClick"]),x.showSelector?(o(),n(v,{key:0,class:"uni-select--mask",onClick:k.toggleSelector},null,8,["onClick"])):d("",!0),x.showSelector?(o(),n(v,{key:1,class:"uni-select__selector"},{default:r((()=>[m(v,{class:"uni-popper__arrow"}),m(R,{"scroll-y":"true",class:"uni-select__selector-scroll"},{default:r((()=>[0===x.mixinDatacomResData.length?(o(),n(v,{key:0,class:"uni-select__selector-empty"},{default:r((()=>[m(w,null,{default:r((()=>[f(u(a.emptyTips),1)])),_:1})])),_:1})):(o(!0),h(_,{key:1},b(x.mixinDatacomResData,((e,t)=>(o(),n(v,{class:"uni-select__selector-item",key:t,onClick:t=>k.change(e)},{default:r((()=>[m(w,{class:p({"uni-select__selector__disabled":e.disable})},{default:r((()=>[f(u(k.formatItemName(e)),1)])),_:2},1032,["class"])])),_:2},1032,["onClick"])))),128))])),_:1})])),_:1})):d("",!0)])),_:1},8,["class"])])),_:1},8,["class"])])),_:1})}],["__scopeId","data-v-32096ab0"]]);export{x as _};
