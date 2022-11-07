import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Login from "../views/login.vue";
import Admin from "../views/Admin/admin.vue";
import Addemployee from "../views/Admin/addEmployee.vue"
import HR from "../views/HR/hr.vue";
import Employee from "../views/Employee/employee.vue";
import EmployeeList from "../views/Admin/employeeList.vue"
import HrEmployeeList from "../views/HR/ListEmployees.vue"
import EmployeeProfile from '../views/Employee/profile.vue'
import Approvals from '../views/Employee/approvals.vue'
import ApplyLeaves from '../views/ApplyLeaves.vue'
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
    path: "/admin",
    name: "Admin",
    component: Admin,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/admin/add/employee",
    name: "Addemployee",
    component: Addemployee,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/hr",
    name: "HR",
    component: HR,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/employee",
    name: "Employee",
    component: Employee,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/admin/all/employees",
    name: "EmployeeList",
    component: EmployeeList,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/hr/all/employees",
    name: "HrEmployeeList",
    component: HrEmployeeList,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/employee/profile",
    name: "EmployeeProfile",
    component: EmployeeProfile,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/apply/leaves",
    name: "ApplyLeaves",
    component: ApplyLeaves,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/approvals",
    name: "Approvals",
    component: Approvals,
    meta: {
      requiresAuth: true
  }
  },
  {
    path: "/attendance/record",
    name: "AttendanceRecord",
    component: AttendanceRecord,
    meta: {
      requiresAuth: true
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

 if (requiresAuth && ! authenticatedUser) {
   alert("you need to login first")
   localStorage.clear();
   next('/login')
   location.reload()
 }else{
    next();
 }

});

export default router;





