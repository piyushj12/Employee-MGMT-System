<template>
    <div class="container-fluid">
        <div class="row mt-3">
            <div class="col-sm-12">
                <h3 class="text-center">
                    Leaves Approval
                </h3>
            </div>
        </div>
        <div class="row m-3">
            <div class="col-sm-12">
                 <div class="card">
                    <table class="table table-bordered table-striped">
                        <thead>
                            <tr>
                                <th>Applied on</th>
                                <th>Employee Name</th>
                                 <th>Employee email</th>
                                <th>From</th>
                                <th>To</th>
                                <th>Reamining Holidays</th>
                                <th>No. of days</th>
                                <th>Status</th>                     
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(leaves,index) in appliedleaves" :key="index">
                                <td>{{new Date(leaves.applieddate).toLocaleString()}}</td>
                                <td>{{leaves.details.firstname +" "+ leaves.details.lastname}}</td>
                                <td>{{leaves.email}}</td>
                                <td>{{new Date(leaves.fromdate).toLocaleDateString()}}</td>
                                <td>{{new Date(leaves.todate).toLocaleDateString()}}</td>
                                <td>{{leaves.details.remholidays}}</td>
                                <td>{{leaves.numdays}}</td>
                                <td>{{leaves.status}}</td>
                                <td class="text-center">
                                    <button class="btn btn-primary btn-sm" @click="updateStatus(leaves.lid,leaves.email,'approved',leaves.numdays,leaves.holidays)">Approve</button>&nbsp;
                                    <button class="btn btn-secondary btn-sm">View Previous</button>&nbsp;
                                    <button class="btn btn-danger btn-sm" @click="updateStatus(leaves.lid,leaves.email,'cancelled',leaves.numdays,leaves.holidays)">Cancel</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import Services from '../../services/EmployeeServices.js'
export default {
    data(){
        return{
            appliedleaves:[]
        }
    },
    created(){
        this. getEmployeeAppliedLeaves()
    },
    methods:{
        async getEmployeeAppliedLeaves(){
            await Services.getEmployeeAppliedLeaves({"email":localStorage.getItem('email'),"status":"pending"})
            .then((data) => {
                console.log(data)
                this.appliedleaves=data
            })
        },
        async updateStatus(id,email,status,numdays,holidays){
            await Services.updateLeaves({"lid":id,"email":email,"status":status,"numdays":numdays,"holidays":holidays})
            .then((data) => {
                this.$toast.open({
                    message:data.message,
                    type: "success",
                    position: "top",
                });
                location.reload()
                
            })
        },
    }
}
</script>