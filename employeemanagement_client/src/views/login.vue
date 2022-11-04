<template>
  <div>
    <div class="container-fluid">
      <div class="row">
        <div
          class="col-sm-12 col-lg-8 col-md-8 text-center d-flex justify-content-center p-5"
        >
          <div class="innercontent">
            <h2>Login to Your Account</h2>
            <label class="text-secondary">Login using social networks</label>
            <div class="d-flex mt-2 justify-content-center">
              <span class="bg-primary text-light p-2 ps-3 pe-3 rounded-circle">
                <i class="fab fa-facebook-f fa-lg"></i>
              </span>
              <span class="bg-danger text-light p-2 rounded-circle ms-3">
                <i class="fab fa-google-plus-g fa-lg"></i>
              </span>

              <span class="bg-info text-light p-2 rounded-circle ms-3">
                <i class="fab fa-twitter fa-lg"></i>
              </span>
            </div>
            <h5 class="mt-2 text-secondary">or</h5>
            <div class="signupbox p-3">
              <div class="form-floating mb-3">
                <input
                  v-model="login.email"
                  type="email"
                  class="form-control rounded-pill"
                  id="floatingInput"
                  placeholder="Email Address"
                  :class="{
                            'is-invalid':
                              submitted && $v.login.email.$error,
                          }"
                />
                <div
                          v-if="submitted && $v.login.email.$error"
                          class="invalid-feedback"
                        >
                          <span v-if="!$v.login.email.required"
                            >Email is required</span
                          >
                          <span v-if="!$v.login.email.email"
                            >Must be valid Email</span
                          >
                        </div>
                <label for="floatingInput">Email address</label>
              </div>
              <div class="form-floating">
                <input
                  v-model="login.password"
                  :type="passwordtoggle"
                  class="form-control rounded-pill"
                  id="floatingPassword"
                  placeholder="Password"
                   :class="{
                            'is-invalid':
                              submitted && $v.login.password.$error,
                          }"
                />
                <div
                          v-if="submitted && !$v.login.password.required"
                          class="invalid-feedback"
                        >
                          Password is required
                        </div>
                <label for="floatingPassword">Password</label>
              </div>
              <div class="text-start ms-3">
                <input
                  type="checkbox"
                  @change="PasswordToogle()"
                  v-model="checked"
                  required
                />&nbsp;Show Password
              </div>
              <button class="thmbtn1 w-50 mt-3" @click="Login()">Login</button>
            </div>
          </div>
        </div>
        <div
          class="col-sm-12 col-md-4 col-lg-4 registerpart d-flex text-center flex-column mx-auto justify-content-center align-items-center"
        >
          <div class="p-3">
            <h2 class="text-light">New Here?</h2>
            <h6 class="text-light mt-3">
              Signup and explore more Features.
            </h6>
            <router-link to="/registration" class=" btn btnsignup mt-3">Sign Up</router-link>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import UserServices from "@/services/userService.js";
import { required,email } from "vuelidate/lib/validators";
import VueJwtDecode from 'vue-jwt-decode'

export default {
  data() {
    return {
      submitted:false,
      login: {
        email:"",
        password: "",
      },
      checked: false,
      passwordtoggle: "password",
    };
  },
  validations:{
    login: {
        email:{required,email},
        password: {required},
      }
  },
  methods: {
    PasswordToogle() {
      if (this.checked) {
        this.passwordtoggle = "text";
      } else {
        this.passwordtoggle = "password";
      }
    },
    async Login() {
       this.submitted = true;
      this.$v.login.$touch();
      if (this.$v.login.$invalid) {
        return;
      }
      this.login.empstatus="active";
      await UserServices.loginService(this.login).then((data) => {
          if(data.Error!=null){
            this.$toast.open({
            message:data.Error,
            type: "error",
            position: "top",
          });
          }else{
            //console.log(VueJwtDecode.decode(data.Authorization))
              localStorage.setItem('token',data.Authorization)
              localStorage.setItem('role',VueJwtDecode.decode(data.Authorization).role)
              localStorage.setItem('email',VueJwtDecode.decode(data.Authorization).email)
              if(localStorage.getItem('role')=='admin'){
                this.$router.push({ path: "/admin" });
              }else if(localStorage.getItem('role')=='hr'){
                this.$router.push({ path: "/hr" });
              }else{
                this.$router.push({ path: "/employee" });
              }
              location.reload();
          }
      });
    },
  },
};
</script>
<style scoped src="../assets/CSS/login.css">

</style>
