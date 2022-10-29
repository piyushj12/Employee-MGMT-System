<template>
  <div class="container-fluid">
    <div class="row mt-3">
        <div class="col-sm-12">
            <h3 class="text-center"> Employees List Admin</h3>
        </div>
    </div>
    <div class="row mt-3">
        <div class="col-sm-12">
             <EmployeeList/>
        </div>
    </div>
  </div>
</template>

<script>
import EmployeeList from "../../components/EmployeeList.vue"
export default {
    components: {
        EmployeeList:EmployeeList

  },
    data(){
        return{
          
        }
    },
}
</script>

<style>

</style>