<template>
    <div class="container-fluid">
        <div class="row mt-3">
            <div class="col-sm-12">
                    
            </div>
        </div>
        <div class="row mt-3 justify-content-center">
            <div class="col-sm-6">
                 <h3 class="text-center">Apply for Leave</h3>
                <div class="card p-5">
                   <label class="labels">Number of leaves remaining:</label>
                   <label >{{leaves.remholidays}}</label><br/>
                   <label class="labels">Approval required by:</label>
                   <label >{{leaves.manageremail}}</label><br/>
                    <label class="labels">Select dates:</label>
                    <date-picker @change="countDays" v-model="dates" value-type="format" range></date-picker><br/>
                    <label class="labels">Number of days selected:</label>
                    <input type="text" v-model="days" disabled class="form-control" placeholder="Number of days"/><br/>
                    <label class="labels">Description:</label>
                    <textarea v-model="comment" class="form-control"></textarea><br/>
                    <button class="btn btn-primary w-50" @click="submitforapproval">Submit for Approval</button>
                </div>
            </div>
            <div class="col-sm-6">
                <h3 class="text-center">Applied Leaves</h3>
                <AppliedLeaves :email="mail" status="applied" field="email" />
            </div>
        </div>
    </div>
</template>
<script>
import Services from '../services/EmployeeServices.js'
import DatePicker from 'vue2-datepicker';
import 'vue2-datepicker/index.css';
import AppliedLeaves from '@/components/appliedLeaves.vue';
export default {
     components: { DatePicker,AppliedLeaves},
    data(){
        return{
            leaves:{},
            dates:[],
            days:"",
            appliedleaves:[],
            comment:"",
            mail:""
        }
    },
    created(){
        this.mail=localStorage.getItem('email')
        this.getEmployeeLeaves()
        this.getEmployeeAppliedLeaves()
    },
    methods:{
        countDays(){
             var Difference_In_Time  = new Date(this.dates[1]).getTime() - new Date(this.dates[0]).getTime();
            this.days = Difference_In_Time / (1000 * 3600 * 24) + 1;
        },
        async getEmployeeLeaves(){
            await Services.getEmployeeLeaves({"email":localStorage.getItem('email')})
            .then((data) => {
              
                 this.leaves=data[0]
            })
        },
         async getEmployeeAppliedLeaves(){
            await Services.getEmployeeAppliedLeaves({"email":localStorage.getItem('email'),"status":"applied"})
            .then((data) => {
                 if(data!=null){
                    this.appliedleaves=data
                 }else{
                    this.appliedleaves=[]
                 }
                
            })
        },
        async submitforapproval(){
            let obj={}
            obj.email=localStorage.getItem('email')
            obj.holidays=this.leaves.holidays
            obj.numdays=this.days
            obj.manageremail=this.leaves.manageremail
            obj.comment=this.comment
            obj.fromdate= new Date(this.dates[0].split("-")[0],(parseInt(this.dates[0].split("-")[1])-1).toString(),this.dates[0].split("-")[2]).getTime()
            obj.todate=  new Date(this.dates[1].split("-")[0],(parseInt(this.dates[1].split("-")[1])-1).toString(),this.dates[1].split("-")[2]).getTime()
            await Services.submitforapproval(obj)
            .then((data) => {
                if(data){
                     this.$toast.open({
                    message:"Applied successfully",
                    type: "success",
                    position: "top",
                });
                location.reload()

                }else{
                      this.$toast.open({
                    message:"Something went wrongPlease contact your manager",
                    type: "error",
                    position: "top",
                });
                }
                
            })
        }
    }
}
</script>

<style>
.labels{
    font-weight: 600;
    font-size: 20px;
}
</style>