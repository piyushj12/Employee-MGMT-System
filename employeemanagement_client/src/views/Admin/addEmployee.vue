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
                <label>Gender:</label>
                <select v-model="emp.gender" class="form-control">
                     <option disabled value="">Please Select Role</option>
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                    <option value="Other">Other</option>
                </select>
                 <label>Department Name:</label>
                     <select  v-model="emp.department" class="form-control">
                       <option disabled value="">Please Select Department</option>
                    <option value="Administration">Administration</option>
                    <option value="Technical support">Technical support</option>
                    <option value="Application development">Application development</option>
                    <option value="IT security">IT security</option>
                    <option value="Human Resource">Human Resource</option>
                </select>
                 <label>Address:</label>
                <input type="text" placeholder="Enter employee's address" v-model="emp.address" class="form-control"/>

                 <label>Email:</label>
                <input type="text" placeholder="Enter employee's email" v-model="emp.email" class="form-control"/>

                 <label>Contact:</label>
                <input type="text" placeholder="Enter employee's contact" v-model="emp.contact" class="form-control"/>

                 <label>Level:</label>
                     <select v-model="emp.level" class="form-control">
                       <option disabled value="">Please Select Level</option>
                    <option value="New Graduate">New Graduate</option>
                    <option value="SDE 1">SDE 1</option>
                    <option value="SDE 2">SDE 2</option>
                    <option value="SDE 3">SDE 3</option>
                    <option value="SDE 4">SDE 4</option>
                    <option value="Project Manager">Project Manager</option>
                   <option value="Recruiting Manager">Recruiting Manager</option>
                   <option value="Managing Director">Managing Director</option>
                    <option value="CEO">CEO</option>
                     <option value="CFO">CFO</option>
                      <option value="CTO">CTO</option>
                </select>
                 <label>Role:</label>
                <select v-model="emp.role" class="form-control">
                     <option disabled value="">Please Select Role</option>
                    <option value="hr">HR</option>
                    <option value="employee">Employee</option>
                </select>
                 <label>Manager:</label>
                <select v-model="emp.manageremail" class="form-control" @change="setManager($event)">
                     <option disabled value="">Please Select Manager</option>
                    <option v-for="(manager,index) in managers" :key="index" :value="manager.email">
                            {{getTitleCase(manager.firstname + " " + manager.lastname) + " " +"("+(manager.department)+")"}}
                    </option>
                </select>
                <label>Yearly Holidays:</label>
                 <select v-model.number="emp.holidays" class="form-control">
                     <option disabled value="">Please Select</option>
                    <option value="10">10</option>
                    <option value="20">20</option>
                    <option value="30">30</option>
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
                department:"",
                manager:"",
                manageremail:"",
                holidays:"",
                gender:""

            },
            managers:[]
        }
    },
    created(){
        this.getManagers()
    },
    methods:{
        async getManagers(){
             await UserServices.getManagerService().then((data) => {
                   this.managers=data
            });
        },
        setManager(event){
                this.emp.manageremail=event.target.value
        },
        async addEmployee(){
             await UserServices.addEmployeeService(this.emp).then((data) => {
           this.$toast.open({
            message:data.message,
            type: "success",
            position: "top",
          });
          this.emp={}
            this.getManagers()
          //console.log(data)
      });
            
        }
    }
}
</script>

<style>

</style>