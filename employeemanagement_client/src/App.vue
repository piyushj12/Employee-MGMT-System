<template>
 <div id="app" class="d-flex flex-column min-vh-100 scrollclass">
      <headerbar/>
      <div class="wrapper">
           <navigationbar v-if="token!=null" class="first"/>
            <router-view />
      </div>
     
      <footerbar/>
  </div>
</template>
<script>
import headerbar from "./components/header.vue";
import footerbar from "./components/footer.vue";
import navigationbar from "./components/navigationbar.vue";
export default{
  data(){
    return{
       token:"djdj",
       sidenav:"sidenav"
    }
  },
  created(){
      this.token=localStorage.getItem('token')
      if(this.token==null){
          this.sidenav="sidenav"
      }else{
         this.sidenav="sidenavlogin"
      }
  },
  name: "app",
  components: {
    /* eslint-disable vue/no-unused-components */
    headerbar: headerbar,
     /* eslint-disable vue/no-unused-components */
    footerbar:footerbar,
    navigationbar:navigationbar

  },
}
</script>

<style lang="scss">
:root {
	--primary: #4ade80;
	--primary-alt: #22c55e;
	--grey: #64748b;
	--dark: #1e293b;
	--dark-alt: #334155;
	--light: #f1f5f9;
	--sidebar-width: 200px;
}
* {
	margin: 0;
	padding: 0;
	box-sizing: border-box;
	font-family: 'Fira sans', sans-serif;
}
body {
	background: var(--light);
  overflow-x: hidden;
}
button {
	cursor: pointer;
	appearance: none;
	border: none;
	outline: none;
	background: none;
}
.app {
	main {
		flex: 1 1 0;
		padding: 2rem;
		@media (max-width: 1024px) {
			padding-left: 6rem;
		}
	}
}
.sidenavlogin{
  display: grid;
  grid-template-columns:repeat(2, 1fr);
  grid-auto-columns: 100px 800px;
}
.sidenav{
   display: grid;
  grid-template-columns: 1fr;
}

.wrapper{
    width: 100vw;
}
.first{
  //  width: 300px;
  float: left;
}
.second{
  margin-left: 305px;
}
</style>