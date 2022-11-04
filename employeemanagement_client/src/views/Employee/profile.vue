<template>
    <div class="container-fluid">
    <b-modal ref="modal-editprofile" title="Edit Profile">
      <div class="p-3 d-flex flex-column justify-content-center">
       <div class="form-group">
        <table class="table table-bordered table-striped">
            <tbody>
                <tr>
                    <th>Address</th>
                    <td>
                        <textarea class="form-control" v-model="profile.address"></textarea>
                    </td>
                </tr>
                 <tr>
                    <th>Contact</th>
                    <td>
                       <input type="text" class="form-control" v-model="profile.contact" />
                    </td>
                </tr>
                 <tr>
                    <th>Password</th>
                    <td>
                        <input type="password" class="form-control" v-model="profile.password" />
                    </td>
                </tr>
            </tbody>
        </table>
        </div>
      </div>
      <template #modal-footer="{ cancel }">
        <b-button size="md" variant="secondary" @click="cancel()">
          Cancel
        </b-button>
        <button type="button" class="btn btn-primary" @click="EditProfile()">
          Edit
        </button>
      </template>
    </b-modal>
        <div class="row mt-3">
            <div class="col-sm-12">
                <h3 class="text-center">Profile</h3>
            </div>
        </div>
        <div class="row mt-3 justify-content-center p-2">
            <div class="col-sm-8">
                <div class="card">
                    <div class="text-center">
                        <img src="../../assets/Images/avatar.png" height="150" width="150" alt="profile picture"/>
                    </div>
                    <table class="table table-bordered table-striped">
                        <tbody>
                            <tr>
                                <th>Name</th>
                                <td>{{profile.firstname+" "+profile.lastname}}</td>
                            </tr>
                             <tr>
                                <th>Email</th>
                                <td>{{profile.email}}</td>
                            </tr>
                             <tr>
                                <th>Contact</th>
                                <td>{{profile.contact}}</td>
                            </tr>
                             <tr>
                                <th>Address</th>
                                <td>{{profile.address}}</td>
                            </tr>
                             <tr>
                                <th>Department</th>
                                <td>{{profile.department}}</td>
                            </tr>
                             <tr>
                                <th>Position</th>
                                <td>{{profile.level}}</td>
                            </tr>
                             <tr>
                                <th>Manager Email</th>
                                <td>{{profile.manageremail}}</td>
                            </tr>
                             <tr>
                                <th>Holidays</th>
                                <td>{{profile.holidays}}</td>
                            </tr>
                            <tr>
                                <th>Password</th>
                                <td>{{profile.password}}</td>
                            </tr>
                        </tbody>
                    </table>
                    <div class="text-center pb-2">
                        <button class="btn btn-primary" @click="openEditModel()">Edit Profile</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>
<script>
import Services from "@/services/EmployeeServices.js";
export default {
   data(){
    return{
        profile:{},
    }
   },
   created(){
        this.getEmployee()
   },
   methods:{
        async getEmployee(){
            await Services.getProfile({"email":localStorage.getItem('email')})
            .then((data) => {
                this.profile=data[0]
            })
        },
        async EditProfile(){
             await Services.editProfile({"email":this.profile.email,"address":this.profile.address,"contact":this.profile.contact,"password":this.profile.password})
            .then(() => {
                this.getEmployee()
                this.$refs["modal-editprofile"].hide();
            })
        },
        openEditModel(){
             this.$refs["modal-editprofile"].show();

        }
   }
}
</script>