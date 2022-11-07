<template>
    <div class="container-fluid">
        <b-modal ref="modal-view" size="lg" title="Previous Leave Records">
      <div class="p-3 d-flex flex-column justify-content-center">
        <div class="form-group">
            <div class="text-end">
                <button class="btn btn-success btn-md" @click="setStatus('approved')">Approved Leaves</button>&nbsp;
                <button class="btn btn-danger btn-md"  @click="setStatus('cancelled')">Cancelled Leaves</button>
            </div>
            
            <AppliedLeaves :email="email" :status="status" field="email" :key="status" class="mt-3"/>

        </div>
      </div>
      <template #modal-footer="{ cancel }">
        <b-button size="md" variant="secondary" @click="cancel()">
          Cancel
        </b-button>
      </template>
    </b-modal>
        <div class="row mt-3">
            <div class="col-sm-12">
                <h3 class="text-center">
                    Leaves Approval
                </h3>
            </div>
        </div>
        <div class="row m-3">
            <div class="col-sm-12">
                 <div class="card" v-if="appliedleaves.length>0">
                    <table class="table table-bordered table-striped">
                        <thead>
                            <tr>
                                <th>Applied on</th>
                                <th>Employee Name</th>
                                 <th>Employee email</th>
                                <th>From</th>
                                <th>To</th>
                                <th>Rem. Holidays</th>
                                <th>No. of days</th>
                                <th>Status</th>                     
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(leaves,index) in appliedleaves" :key="index">
                                <td>{{new Date(leaves.applieddate).toLocaleDateString()}}</td>
                                <td>{{leaves.details.firstname +" "+ leaves.details.lastname}}</td>
                                <td>{{leaves.email}}</td>
                                <td>{{new Date(leaves.fromdate).toLocaleDateString()}}</td>
                                <td>{{new Date(leaves.todate).toLocaleDateString()}}</td>
                                <td>{{leaves.details.remholidays}}</td>
                                <td>{{leaves.numdays}}</td>
                                <td>{{leaves.status}}</td>
                                <td class="" style="min-width:240px">
                                    <button class="btn btn-primary btn-sm" @click="updateStatus(leaves.lid,leaves.email,'approved',leaves.numdays,leaves.holidays)">Approve</button>&nbsp;
                                    <button class="btn btn-secondary btn-sm" @click="viewRecordsModal(leaves.email,'approved')">Previous</button>&nbsp;
                                    <button class="btn btn-danger btn-sm" @click="updateStatus(leaves.lid,leaves.email,'cancelled',leaves.numdays,leaves.holidays)">Cancel</button>
                                </td>
                            </tr>
                        </tbody>
                    </table>
                </div>
                <NoData v-else/>
            </div>
        </div>
    </div>
</template>
<script>
import Services from '../../services/EmployeeServices.js'
import NoData from '@/components/noDataTemplate.vue';
import AppliedLeaves from '@/components/appliedLeaves.vue';
export default {
    components: {
  NoData,
  AppliedLeaves
},
    data(){
        return{
            appliedleaves:[],
            email:"",
            status:"",
        }
    },
    created(){
        this. getEmployeeAppliedLeaves()
    },
    methods:{
        async getEmployeeAppliedLeaves(){
            await Services.getEmployeeAppliedLeaves({"email":localStorage.getItem('email'),"status":"pending","field":"manageremail"})
            .then((data) => {
               if(data!=null){
                     this.appliedleaves=data
               }else{
                     this.appliedleaves=[]
               }
               
            })
        },
        setStatus(status){
            this.status=status
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
        viewRecordsModal(email,status){
            this.email=email
            this.status=status
            this.$refs["modal-view"].show();

        }
    }
}
</script>