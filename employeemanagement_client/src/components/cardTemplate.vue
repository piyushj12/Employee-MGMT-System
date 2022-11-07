<template>
  <div :class="containerClass">
    <div class="text-light">
      <h1 class="iconpos">{{value}}</h1>
      <h3 class="pt-4 ps-3">{{title}}</h3>
    </div>
  </div>
</template>
<script>
export default {
    props: {
    title: String,
    color:String,
    icon:String,
    value:String
  },
  data(){
    return{
        iconClass:"iconpos fa-7x",
        containerClass:"container-fluid customcard"
    }
  },
  created(){
    this.iconClass= this.iconClass+" "+this.icon
    this.containerClass=this.containerClass+" "+this.color
    console.log()
  }
}
</script>
<style>
.customcard {
  height: 200px;
  width: 95%;
  border-radius: 20px;
  position: relative;
  cursor: pointer;
}
.card-color-green {
  background: rgb(31, 158, 62);
  background: linear-gradient(
    90deg,
    rgba(31, 158, 62, 1) 19%,
    rgba(3, 169, 125, 0.6647701316854867) 100%
  );
}
.card-color-tomato {
  background: rgb(233, 61, 35);
  background: linear-gradient(
    90deg,
    rgba(233, 61, 35, 0.8076272745426296) 24%,
    rgba(185, 54, 22, 0.5331174706210609) 100%
  );
}
.card-color-yellow {
  background: rgb(236, 195, 46);
  background: linear-gradient(
    90deg,
    rgba(236, 195, 46, 0.8076272745426296) 24%,
    rgba(185, 172, 22, 0.5331174706210609) 100%
  );
}
.card-color-blue {
  background: rgb(46, 120, 236);
  background: linear-gradient(
    90deg,
    rgba(46, 120, 236, 0.8076272745426296) 24%,
    rgba(22, 141, 185, 0.45188497762386204) 100%
  );
}

.card-color-purple {
  background: rgb(94, 6, 89);
  background: linear-gradient(
    90deg,
    rgba(94, 6, 89, 0.9644900196406687) 24%,
    rgba(104, 29, 129, 0.863649683506215) 100%
  );
}

.card-color-darkgreen {
  background: rgb(6, 94, 34);
  background: linear-gradient(
    90deg,
    rgba(6, 94, 34, 0.9364788151588761) 24%,
    rgba(25, 127, 42, 0.7263947815454307) 100%
  );
}
.iconpos {
  position: absolute;
  bottom: 0;
  padding-bottom: 10px;
  right: 0;
  padding-right: 20px;
  opacity: 0.75;
  font-size: 35px;
  padding-left:20px
}
@media only screen and (max-width: 1024px) {
  .iconpos {
    font-size: 25px;
  }
}
</style>
