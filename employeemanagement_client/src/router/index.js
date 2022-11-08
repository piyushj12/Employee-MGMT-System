import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/login.vue";
import Addemployee from "../views/Admin/addEmployee.vue"
import EmployeeList from "../views/Admin/employeeList.vue"
import Settings from "../views/Admin/settings.vue"
import HrEmployeeList from "../views/HR/ListEmployees.vue"
import EmployeeProfile from '../views/Employee/profile.vue'
import Approvals from '../views/Employee/approvals.vue'
import ApplyLeaves from '../views/ApplyLeaves.vue'
import Dashboard from '../views/dashboard.vue'
import AttendanceRecord from '../views/attendanceRecord.vue'
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
  {
    path: "/login",
    name: "Login",
    component: Login,
  },
  {
    path: "/login/dashboard",
    name: "Dashboard",
    component: Dashboard,
    meta: {
      requiresAuth: true,
  }
  },
  {
    path: "/admin/add/employee",
    name: "Addemployee",
    component: Addemployee,
    meta: {
      requiresAuth: true,
      adminAuth:true,
      empAuth:false,
      hrAuth:true
  }
  },
 
  {
    path: "/admin/all/employees",
    name: "EmployeeList",
    component: EmployeeList,
    meta: {
      requiresAuth: true,
      adminAuth:true,
      empAuth:false,
      hrAuth:false
  }
  },
  {
    path: "/hr/all/employees",
    name: "HrEmployeeList",
    component: HrEmployeeList,
    meta: {
      requiresAuth: true,
      adminAuth:false,
      empAuth:false,
      hrAuth:true
  }
  },
  {
    path: "/employee/profile",
    name: "EmployeeProfile",
    component: EmployeeProfile,
    meta: {
      requiresAuth: true,
      
  }
  },
  {
    path: "/apply/leaves",
    name: "ApplyLeaves",
    component: ApplyLeaves,
    meta: {
      requiresAuth: true,
      adminAuth:false,
      empAuth:true,
      hrAuth:true

  }
  },
  {
    path: "/approvals",
    name: "Approvals",
    component: Approvals,
    meta: {
      requiresAuth: true,
     }
  },
  {
    path: "/attendance/record",
    name: "AttendanceRecord",
    component: AttendanceRecord,
    meta: {
      requiresAuth: true,
      adminAuth:false,
      empAuth:true,
      hrAuth:true
  }
  },
  {
    path: "/settings",
    name: "Settings",
    component: Settings,
    meta: {
      requiresAuth: true,
      adminAuth:true,
      empAuth:false,
      hrAuth:false
  }
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

router.beforeEach((to, from, next) => {
  if(to.name=='Login'){
    if(localStorage.getItem('token')!=null){
      alert("You are already logged in")
    }else{
      next()
    }
  }
 var authenticatedUser = null;
 if(localStorage.getItem('token')!=null){
    authenticatedUser=true;
 }

 const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

 if (requiresAuth && !authenticatedUser) {
   alert("you need to login first")
   localStorage.clear();
   next('/login')
   location.reload()
 }else{
  if(to.meta.adminAuth && to.meta.hrAuth){
    const role=localStorage.getItem('role')
    if(role==="admin" || role==="hr"){
      next()
    }else{
      alert("Unauthorized access")
      next('/login/dashboard')
    }
  }else if(to.meta.hrAuth && to.meta.empAuth){
    const role=localStorage.getItem('role')
    if(role==="employee" || role==="hr"){
      next()
    }else{
      alert("Unauthorized access")
      next('/login/dashboard')
    }
  }else if(to.meta.adminAuth){
      const role=localStorage.getItem('role')
      if(role==="admin"){
        next()
      }else{
        alert("Unauthorized access")
        next('/login/dashboard')
      }
    }else if(to.meta.empAuth){
      const role=localStorage.getItem('role')
      if(role==="employee"){
        next()
      }else{
        alert("Unauthorized access")
        next('/login/dashboard')
      }
    }else if(to.meta.hrAuth){
      const role=localStorage.getItem('role')
      if(role==="hr"){
        next()
      }else{
        alert("Unauthorized access")
        next('/login/dashboard')
      }
    }else{
      next();
    }
   
 }

});

export default router;





