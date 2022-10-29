<template>
  <footer class="container-fluid footer mt-auto">
    <div class="row pt-3 pb-3 justify-content-center">
      <div class="col-sm-12 col-md-6 col-lg-4">
        <div class="ms-5 mt-3 footerlogomobile">
          <router-link to="/" class="text-light anchorstyle">
            <img
              src="../assets/Images/logo.png"
              height="100"
              width="100"
              class="ms-4 mobilefooterlogo"
              alt="logo"
            />
            <h4 class="mobilefooterlogotext">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;EMS</h4>
          </router-link>
        </div>
      </div>
      <div class="col-sm-12 col-md-6 col-lg-4 text-left">
        <div class="ms-5 mt-2">
         <router-link class="text-light linkclass" to="/contact"><h5>Contact</h5></router-link> 
          <i class="fa fa-home fa-lg mt-3"></i>&nbsp;8765 University Road Apt L
          <br />
          <i class="fas fa-phone fa-lg pt-4"></i>&nbsp;(980) 327 9722
          <br />
          <a class="text-light anchorstyle" href="mailto:abc@example.com'">
            <i class="fas fa-envelope fa-lg pt-4 pb-2"></i
            >&nbsp;tnemade@uncc.edu
          </a>
        </div>
      </div>
      <div class="col-sm-12 col-md-6 col-lg-4">
        <div class="ms-5 mt-2">
          <h5 class="">Follow Us</h5>
          <div class="mt-3 text-light">
            <!-- <a href="#">
              <i class="fab fa-facebook fa-lg pe-2 text-light"></i
            ></a> -->
            <a href="#"><i class="fab fa-linkedin fa-lg p-2 text-light"></i></a>
            <a href="#">
              <i class="fab fa-twitter-square fa-lg p-2 text-light"></i
            ></a>
            <a href="#">
              <i class="fab fa-instagram fa-lg p-2 text-light"></i
            ></a>
              <a href="#" class="rounded">
              <i class="fa fab fa-facebook-f fa-lg p-2 text-light"></i
            ></a>
           
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-sm-12 text-center text-light p-1 lastfooter">
        <p class="mt-2">
          Copyright&nbsp;<i class="far fa-copyright"></i>&nbsp;2022 Employee Management System
        </p>
      </div>
    </div>
  </footer>
</template>
<style scoped>
@import "../assets/CSS/common.css";
</style>
