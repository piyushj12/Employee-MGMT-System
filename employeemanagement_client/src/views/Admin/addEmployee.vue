<template>
  <div class="container-fluid">
    <div class="row">
        <div class="col-sm-12">
            <h3 class="text-center mt-3">Add Employee</h3>
        </div>
    </div>
        <div class="row p-3 justify-content-center">
            <div class="col-sm-7">
                <div class="card p-5">
                <label>First Name:</label>
                <input type="text" placeholder="Enter employee's first name" v-model="emp.firstname" class="form-control"/>
                <label>Last Name:</label>
                <input type="text" placeholder="Enter employee's last name" v-model="emp.lastname" class="form-control"/>

                 <label>Department Name:</label>
                <input type="text" placeholder="Enter employee's department name" v-model="emp.department" class="form-control"/>

                 <label>Address:</label>
                <input type="text" placeholder="Enter employee's address" v-model="emp.address" class="form-control"/>

                 <label>Email:</label>
                <input type="text" placeholder="Enter employee's email" v-model="emp.email" class="form-control"/>

                 <label>Contact:</label>
                <input type="text" placeholder="Enter employee's contact" v-model="emp.contact" class="form-control"/>

                 <label>Level:</label>
                <input type="text" placeholder="Enter employee's level" v-model="emp.level" class="form-control"/>

                 <label>Role:</label>
                <select v-model="emp.role" class="form-control">
                    <option value="hr">HR</option>
                    <option value="employee">Employee</option>
                </select>
                <button class="btn btn-primary w-50 mt-3" @click="addEmployee()">Add</button>
                </div>
                
            </div>
        </div>
  </div>
</template>

<script>
import UserServices from "@/services/userService.js";
export default {
    name:"Addemployee",
    data(){
        return{
            emp:{
                firstname:"",
                lastname:"",
                contact:"",
                email:"",
                address:"",
                role:"",
                level:"",
                department:""
            }
        }
    },
    methods:{
        async addEmployee(){
             await UserServices.addEmployeeService(this.emp).then((data) => {
           this.$toast.open({
            message:data.message,
            type: "success",
            position: "top",
          });
          this.emp={}
          //console.log(data)
      });
            
        }
    }
}
</script>

<style>

</style>