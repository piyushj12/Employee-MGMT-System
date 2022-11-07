<template>
    <div class="container-fluid">
        <div class="row mt-3">
            <div class="col-sm-12">
                <h3 class="text-center">Attendance record</h3>
            </div>
        </div>
         <div class="row mt-3">
            <div class="col-sm-12">
                 <label class="mt-3">Select Year</label>
                <select v-model.number="year" class="form-control">
                     <option disabled value="">Please select year</option>
                    <option v-for="(year,index) in years" :value="year" :key="index">{{year}}</option>
                </select>
                <label class="mt-3">Select Month</label>
                <select v-model="month" class="form-control">
                    <option disabled value="">Please select month</option>
                    <option v-for="(month,index) in months" :value="index" :key="index">{{month}}</option>
                </select>
               <button class="btn btn-primary w-25 mt-3" @click="viewAttendanceData()">View</button>
            </div>
        </div>
        <div class="row mt-5">
            <div class="col-sm-12">
                <table  v-if="attendance.length>0" class="table table-bordered table-striped">
                    <thead>
                        <th>Sr. No.</th>
                        <th>Date</th>
                        <th>Clock-In Time</th>
                        <th>Clock-out Time</th>
                        <th>Number of hours worked</th>
                    </thead>
                    <tbody>
                        <tr v-for="(day,index) in attendance" :key="index">
                            <td>{{index+1}}</td>
                            <td>{{new Date(day.clockin).toLocaleDateString()}}</td>
                            <td>{{new Date(day.clockin).toLocaleTimeString()}}</td>
                            <td>
                                <span v-if="day.status">
                                    {{new Date(day.clockout).toLocaleTimeString()}}
                                </span>
                                <span v-else>
                                    N/A
                                </span>
                            </td>
                            <td>
                                 <span v-if="day.status">
                                     {{countDays(day.clockin,day.clockout)}}
                                </span>
                                <span v-else>
                                    N/A
                                </span>
                            </td>
                        </tr>
                    </tbody>
                </table>
                 <NoData  v-else/>
            </div>
        </div>
       
    </div>

</template>
<script>
import Services from '../services/EmployeeServices.js'
import NoData from '@/components/noDataTemplate.vue';
export default {
     components: {
  NoData,
},
    data(){
        return{
            months:["January","February","March","April","May","June","July",
            "August","September","October","November","December"],
            years:["2022"],
            attendance:[],
            year:"",
            month:""
        }
    },
    created(){

    },
    methods:{
                countDays(date1,date2){
             var Difference_In_Time  = new Date(date2).getTime()-new Date(date1).getTime();
            var min=Math.floor((Difference_In_Time /1000)/60);
            var hour=Math.floor(min/60)
            var mins=min%60
            console.log(min)
            return hour+"hr(s)"+" "+mins+"min(s)"
        },
        viewAttendanceData(){
            let obj={}
            obj.month=this.month,
            obj.year=this.year
            obj.empid=localStorage.getItem('empid')
            Services.viewAttendanceData(obj)
            .then(data=>{
                this.attendance=data
            })
        }
    }
}
</script>