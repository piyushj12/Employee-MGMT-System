<template>
    <div class="container-fluid">
        <div class="row mt-3">
            <div class="col-sm-12">
                <h3 class="text-center text-dark">Settings</h3>
            </div>
        </div>
        <div class="row mt-3">
            <div class="col-sm-12">
                <div class="text-end">
                    <button class="btn btn-secondary" @click="ResetEmployeeData()">Yearly reset</button>
                </div>
            </div>
        </div>
        <div class="row mt-3">
            <div class="col-sm-4">
                <div class="card p-5">
                    <h4 class="text-center text-dark">Add Department</h4>
                   <input type="text" placeholder="Enter Deparment Name" class="form-control" v-model="department"/>
                    <button class="btn btn-primary btn-sm mt-3" @click="AddField('departments',department)">Add</button>
                </div>
            </div>
            <div class="col-sm-4">
                <div class="card p-5">
                    <h4 class="text-center text-dark">Add Level</h4>
                   <input type="text" placeholder="Enter Level Name" class="form-control" v-model="level"/>
                   <button class="btn btn-primary btn-sm mt-3" @click="AddField('levels',level)">Add</button>
                </div>
            </div>
            <div class="col-sm-4">
                <div class="card p-5">
                    <h4 class="text-center text-dark">Add Yearly Holidays</h4>
                   <input type="text" placeholder="Enter Level Name" class="form-control" v-model.number="holidays"/>
                   <button class="btn btn-primary btn-sm mt-3" @click="AddField('holidays',holidays)">Add</button>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import Service from '../../services/userService.js'
export default {
    data(){
        return{
            department:"",
            level:"",
            holidays:""
        }
    },
    created(){

    },
    methods:{
        ResetEmployeeData(){
            Service.ResetData()
            .then(data=>{
                if(data){
                        this.$toast.open({
                            message:"Successful",
                            type: "success",
                            position: "top",
                        });
                    }else{
                         this.$toast.open({
                            message:"Something went wrong",
                            type: "error",
                            position: "top",
                        });
                    }
            })
        },
            AddField(field,value){
                let obj={}
                obj.field=field
                obj.value=value
                Service.AddField(obj)
                .then(data=>{
                    if(data){
                        this.$toast.open({
                            message:"Added Successfully",
                            type: "success",
                            position: "top",
                        });
                    }else{
                         this.$toast.open({
                            message:"Something went wrong",
                            type: "error",
                            position: "top",
                        });
                    }
                })
            }
    }
}
</script>