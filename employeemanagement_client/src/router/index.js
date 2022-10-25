import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

// router.beforeEach((to, from, next) => {
//   if(to.name=='Login'){
//     if(localStorage.getItem('id')!=null){
//       alert("You are already logged in")
//       next('/users/loginhome')
//     }else{
//       next()
//     }
//   }
//  var authenticatedUser = null;
//  if(localStorage.getItem('token')!=null){
//    authenticatedUser=true;
//  }

//  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

//  if (requiresAuth && ! authenticatedUser) {
//    alert("you need to login first")
//    localStorage.clear();
//    next('/login')
//    location.reload()
//  }else{
//     next();
//  }

// });

export default router;





