import{_ as e,aG as t,aH as o,o as a,c as n,w as i,i as r,a as s,z as l,A as c,F as f,b as h,r as u,d,e as b,a0 as T,v as g,S as v,t as P}from"./index-886f4693.js";var m=null;const k=e({data:()=>({title:"createContext",names:["rotate","scale","reset","translate","save","restore","drawImage","fillText","fill","stroke","clearRect","beginPath","closePath","moveTo","lineTo","rect","arc","quadraticCurveTo","bezierCurveTo","setFillStyle","setStrokeStyle","setGlobalAlpha","setShadow","setFontSize","setLineCap","setLineJoin","setLineWidth","setMiterLimit"]}),onReady:function(){m=t("canvas",this)},methods:{toTempFilePath:function(){o({canvasId:"canvas",success:function(e){console.log(e.tempFilePath)},fail:function(e){console.error(JSON.stringify(e))}})},handleCanvasButton:function(e){this[e]&&this[e]()},rotate:function(){m.beginPath(),m.rotate(10*Math.PI/180),m.rect(225,75,20,10),m.fill(),m.draw()},scale:function(){m.beginPath(),m.rect(25,25,50,50),m.stroke(),m.scale(2,2),m.beginPath(),m.rect(25,25,50,50),m.stroke(),m.draw()},reset:function(){m.beginPath(),m.setFillStyle("#000000"),m.setStrokeStyle("#000000"),m.setFontSize(10),m.setGlobalAlpha(1),m.setShadow(0,0,0,"rgba(0, 0, 0, 0)"),m.setLineCap("butt"),m.setLineJoin("miter"),m.setLineWidth(1),m.setMiterLimit(10),m.draw()},translate:function(){m.beginPath(),m.rect(10,10,100,50),m.fill(),m.translate(70,70),m.beginPath(),m.fill(),m.draw()},save:function(){m.beginPath(),m.setStrokeStyle("#00ff00"),m.save(),m.scale(2,2),m.setStrokeStyle("#ff0000"),m.rect(0,0,100,100),m.stroke(),m.restore(),m.rect(0,0,50,50),m.stroke(),m.draw()},restore:function(){[3,2,1].forEach((function(e){m.beginPath(),m.save(),m.scale(e,e),m.rect(10,10,100,100),m.stroke(),m.restore()})),m.draw()},drawImage:function(){m.drawImage("../../../static/uni.png",0,0),m.draw()},fillText:function(){m.setStrokeStyle("#ff0000"),m.beginPath(),m.moveTo(0,10),m.lineTo(300,10),m.stroke(),m.setFontSize(10),m.fillText("Hello World",0,30),m.setFontSize(20),m.fillText("Hello World",100,30),m.beginPath(),m.moveTo(0,30),m.lineTo(300,30),m.stroke(),m.draw()},fill:function(){m.beginPath(),m.rect(20,20,150,100),m.setStrokeStyle("#00ff00"),m.fill(),m.draw()},stroke:function(){m.beginPath(),m.moveTo(20,20),m.lineTo(20,100),m.lineTo(70,100),m.setStrokeStyle("#00ff00"),m.stroke(),m.draw()},clearRect:function(){m.setFillStyle("#ff0000"),m.beginPath(),m.rect(0,0,300,150),m.fill(),m.clearRect(20,20,100,50),m.draw()},beginPath:function(){m.beginPath(),m.setLineWidth(5),m.setStrokeStyle("#ff0000"),m.moveTo(0,75),m.lineTo(250,75),m.stroke(),m.beginPath(),m.setStrokeStyle("#0000ff"),m.moveTo(50,0),m.lineTo(150,130),m.stroke(),m.draw()},closePath:function(){m.beginPath(),m.setLineWidth(1),m.moveTo(20,20),m.lineTo(20,100),m.lineTo(70,100),m.closePath(),m.stroke(),m.draw()},moveTo:function(){m.beginPath(),m.moveTo(0,0),m.lineTo(300,150),m.stroke(),m.draw()},lineTo:function(){m.beginPath(),m.moveTo(20,20),m.lineTo(20,100),m.lineTo(70,100),m.stroke(),m.draw()},rect:function(){m.beginPath(),m.rect(20,20,150,100),m.stroke(),m.draw()},arc:function(){m.beginPath(),m.setLineWidth(2),m.arc(75,75,50,0,2*Math.PI,!0),m.moveTo(110,75),m.arc(75,75,35,0,Math.PI,!1),m.moveTo(65,65),m.arc(60,65,5,0,2*Math.PI,!0),m.moveTo(95,65),m.arc(90,65,5,0,2*Math.PI,!0),m.stroke(),m.draw()},quadraticCurveTo:function(){m.beginPath(),m.moveTo(20,20),m.quadraticCurveTo(20,100,200,20),m.stroke(),m.draw()},bezierCurveTo:function(){m.beginPath(),m.moveTo(20,20),m.bezierCurveTo(20,100,200,100,200,20),m.stroke(),m.draw()},setFillStyle:function(){["#fef957","rgb(242,159,63)","rgb(242,117,63)","#e87e51"].forEach((function(e,t){m.setFillStyle(e),m.beginPath(),m.rect(0+75*t,0,50,50),m.fill()})),m.draw()},setStrokeStyle:function(){["#fef957","rgb(242,159,63)","rgb(242,117,63)","#e87e51"].forEach((function(e,t){m.setStrokeStyle(e),m.beginPath(),m.rect(0+75*t,0,50,50),m.stroke()})),m.draw()},setGlobalAlpha:function(){m.setFillStyle("#000000"),[1,.5,.1].forEach((function(e,t){m.setGlobalAlpha(e),m.beginPath(),m.rect(0+75*t,0,50,50),m.fill()})),m.draw(),m.setGlobalAlpha(1)},setShadow:function(){m.beginPath(),m.setShadow(10,10,10,"rgba(0, 0, 0, 199)"),m.rect(10,10,100,100),m.fill(),m.draw()},setFontSize:function(){[10,20,30,40].forEach((function(e,t){m.setFontSize(e),m.fillText("Hello, world",20,20+40*t)})),m.draw()},setLineCap:function(){m.setLineWidth(10),["butt","round","square"].forEach((function(e,t){m.beginPath(),m.setLineCap(e),m.moveTo(20,20+20*t),m.lineTo(100,20+20*t),m.stroke()})),m.draw()},setLineJoin:function(){m.setLineWidth(10),["bevel","round","miter"].forEach((function(e,t){m.beginPath(),m.setLineJoin(e),m.moveTo(20+80*t,20),m.lineTo(100+80*t,50),m.lineTo(20+80*t,100),m.stroke()})),m.draw()},setLineWidth:function(){[2,4,6,8,10].forEach((function(e,t){m.beginPath(),m.setLineWidth(e),m.moveTo(20,20+20*t),m.lineTo(100,20+20*t),m.stroke()})),m.draw()},setMiterLimit:function(){m.setLineWidth(4),[2,4,6,8,10].forEach((function(e,t){m.beginPath(),m.setMiterLimit(e),m.moveTo(20+80*t,20),m.lineTo(100+80*t,50),m.lineTo(20+80*t,100),m.stroke()})),m.draw()}}},[["render",function(e,t,o,m,k,S){const w=u(b("page-head"),d),y=T,L=g,p=v,C=r;return a(),n(C,null,{default:i((()=>[s(w,{title:k.title},null,8,["title"]),s(C,{class:"uni-common-mt"},{default:i((()=>[s(y,{class:"canvas-element","canvas-id":"canvas",id:"canvas"}),s(p,{class:"canvas-buttons","scroll-y":"true"},{default:i((()=>[(a(!0),l(f,null,c(k.names,((e,t)=>(a(),n(L,{key:t,class:"canvas-button",onClick:t=>S.handleCanvasButton(e)},{default:i((()=>[h(P(e),1)])),_:2},1032,["onClick"])))),128)),s(L,{class:"canvas-button",onClick:S.toTempFilePath,type:"primary"},{default:i((()=>[h("toTempFilePath")])),_:1},8,["onClick"])])),_:1})])),_:1})])),_:1})}],["__scopeId","data-v-e3395e83"]]);export{k as default};
