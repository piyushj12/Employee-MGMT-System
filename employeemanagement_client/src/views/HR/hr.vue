<template>
  <div class="container-fluid g-0">
    <div class="row mt-4 g-0">
      <div class="col-sm-12 g-0">
        <h3 class="text-center">Welcome {{ name }}</h3>
      </div>
    </div>
    <div class="row mt-3 p-5">
      <div class="col-sm-6">
        <CardHolder/>
      </div>
      <div class="col-sm-6">
        <Attendance/>
         <SearchEmployee class="mt-3"/>
      </div>
    </div>
  </div>
</template>
<script>
import SearchEmployee from "../../components/searchEmployee.vue";
import CardHolder from "../../components/cardHolder.vue";
import Attendance from "../../components/attendance.vue";
//import Services from "../../services/EmployeeServices.js";
export default {
  components: {
    SearchEmployee,
    CardHolder,
    Attendance
  },
  data() {
    return {
      name: "",
      
    };
  },
  created() {
   
    this.name = localStorage.getItem("email");
  },
  methods: {
  },
};
</script>

<style>
.emp-list{
    cursor: pointer;
   
}
</style>
