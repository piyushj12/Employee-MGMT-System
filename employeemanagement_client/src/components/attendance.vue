<template>
  <div class="container-fluid">
    <div class="card p-4">
      <h4>Attendance</h4>
      <div class="text-center text-secondary mt-2">
        <i class="fa-solid fa-clock fa-3x"></i>
      </div>
      <div class="mt-3 text-center">
        <button class="btn btn-primary w-75" @click="CaptureClockIn()" v-if="!clockInData.status">
            Clock-In
        </button>
        <button class="btn btn-danger w-75" :disabled="clockOutData.status" @click="CaptureClockOut()" v-if="clockInData.status">
            Clock-Out
        </button>
      </div>
      <label class="text-secondary text-center mt-2"
       v-if="clockInData.status" >You clocked in at {{new Date(clockInData.clockin).toLocaleTimeString()}}</label
      >
      <label class="text-secondary text-center mt-2"
       v-if="clockOutData.status" >You clocked out at {{new Date(clockOutData.clockout).toLocaleTimeString()}}</label
      >
    </div>
  </div>
</template>
<script>
import Services from "../services/EmployeeServices.js";
export default {
  data() {
    return {
       clockInData:{
            clockin:"",
            status:false
       },
       clockOutData:{
            clockin:"",
            status:false,
            clockout:""
       },
       btnName:"Clock-In"

    };
  },
  created() {
    this.isClockedIn(),
    this.isClockedOut()
  },
  methods: {
    isClockedIn(){
        Services.isClockedIn({"empid":localStorage.getItem('empid')})
        .then(data=>{
              this.clockInData=data
        })
    },
     isClockedOut(){
        Services.isClockedOut({"empid":localStorage.getItem('empid')})
        .then(data=>{
              this.clockOutData=data
              console.log(this.clockOutData)
        })
    },
    CaptureClockIn() {
      Services.CaptureClockIn({
        email: localStorage.getItem("email"),
        empid: localStorage.getItem("empid"),
      }).then((data) => {
        if (data) {
          this.$toast.open({
            message:" Succesfully Clocked In",
            type: "success",
            position: "top",
          });
          this.isClockedIn()
          this.btnName="Clock-Out"
        }else{
            this.$toast.open({
            message:"Something went wrong contact admin",
            type: "error",
            position: "top",
          });
        }
      });
    },
    CaptureClockOut() {
      Services.CaptureClockOut({
        email: localStorage.getItem("email"),
        empid: localStorage.getItem("empid"),
      }).then((data) => {
        if (data) {
          this.$toast.open({
            message:" Succesfully Clocked Out",
            type: "success",
            position: "top",
          });
            this.isClockedOut()
        }else{
            this.$toast.open({
            message:"Something went wrong contact admin",
            type: "error",
            position: "top",
          });
        }
      });
    },
  },
};
</script>

