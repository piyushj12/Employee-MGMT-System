<template>
  <div class="container-fluid">
    <div class="row mt-3">
        <div class="col-sm-12">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <th>Sr. No.</th>
                        <th>Employee Name</th>
                        <th>Status</th>
                        <th>Email</th>
                        <th>Contact</th>
                         <th>Department</th>
                        <th>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(emp,index) in list" :key="index">
                        <td>{{index+1}}</td>
                        <td>{{emp.firstname+ " "+ emp.lastname }}</td>
                        <td>{{emp.empstatus}}</td>
                        <td>{{emp.email}}</td>
                        <td>{{emp.contact}}</td>
                         <td>{{emp.department}}</td>
                        <td><button class="btn btn-info btn-sm">View</button>&nbsp;<button class="btn btn-secondary btn-sm">Edit</button></td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
  </div>
</template>

<script>
import UserServices from "@/services/userService.js";
export default {
    data(){
        return{
            list:[]
        }
    },
    async created(){
         await UserServices.getAllEmployeesService().then((data) => {
            console.log(data)
                this.list=data;
      });
    }
}
</script>

<style>

</style>