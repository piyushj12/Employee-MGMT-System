<template>
    <div class="container-fluid">
        <div class="row">
          <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Remaining Leaves"
              color="card-color-purple"
              icon="far fa-copyright"
              :value="details.remholidays+'/'+details.holidays"
            />
          </div>
          <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Pending Approvals"
              color="card-color-green"
              icon="far fa-copyright"
              :value="details.approval.pending.toString()"
            />
          </div>
          <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Department"
              color="card-color-tomato"
              icon="far fa-copyright"
              :value="details.department"
            />
          </div>
          <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Manager"
              color="card-color-darkgreen"
              icon="far fa-copyright"
              :value="details.manager"
            />
          </div>
          <!-- <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Check-in Time"
              color="card-color-blue"
              icon="far fa-copyright"
              value="12:00"
            />
          </div>
          <div class="col-sm-12 col-md-12 col-lg-6 p-2">
            <CardTemplate
              title="Check-out Time"
              color="card-color-yellow"
              icon="far fa-copyright"
              value="12:00"
            />
          </div> -->
        </div>
    </div>
</template>
<script>
import CardTemplate from "@/components/cardTemplate.vue";
import Services from "../services/EmployeeServices.js";
export default {
    components: {
    CardTemplate,
  },
  data() {
    return {
      name: "",
      details:{
        holidays:"",
        manager:"",
        remholidays:"",
        approval:{
            pending:""
        }
      }
    };
  },
  created() {
    this.name = localStorage.getItem("email");
    this.DashboardData()
  },
  methods: {
     async DashboardData(){
         await Services.DashboardData({"email":localStorage.getItem('email')}).then((data) => {
        if (data != null) {
            this.details=data
          console.log(data)
        } else {
          console.log(data);
        }
      });
     }
  },
}
</script>