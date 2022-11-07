<template>
  <div class="container-fluid">
    <b-modal ref="modal-view" size="lg" title="Employee Details">
      <div class="">
        <div class="form-group card">
          <div class="row justify-content-center">
            <div class="col-sm-12 text-center">
              <img
                src="../assets/Images/avatar.png"
                height="200"
                width="200"
                alt="profile"
                class=""
              />
            </div>
          </div>
          <div>
            <table class="table table-bordered table-striped">
              <tbody>
                <tr>
                  <th>Name</th>
                  <td>
                    {{ getTitleCase(emp.firstname + " " + emp.lastname) }}
                  </td>
                </tr>
                <tr>
                  <th>Department</th>
                  <td>{{ emp.department }}</td>
                </tr>
                <tr>
                  <th>Email</th>
                  <td>{{ emp.email }}</td>
                </tr>
                <tr>
                  <th>Gender</th>
                  <td>{{ emp.gender }}</td>
                </tr>
                <tr>
                  <th>Contact</th>
                  <td>{{ emp.contact }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
      <template #modal-footer="{ cancel }">
        <b-button size="md" variant="secondary" @click="cancel()">
          Cancel
        </b-button>
      </template>
    </b-modal>
    <div class="">
      <div class="card p-5">
        <div class="row">
          <div class="col-sm-8">
            <h3>Search Employee</h3>
          </div>
          <div class="col-sm-4">
            <div class="text-end">
              <button class="btn btn-primary btn-sm" @click="clearSearch()">
                Clear
              </button>
            </div>
          </div>
        </div>

        <div class="row">
          <div class="col-sm-12 col-md-12 col-lg-5">
            <select class="form-control" v-model="field">
              <option disabled value="">Please Select Field</option>
              <option
                v-for="(field, index) in fields"
                :key="index"
                class="form-control"
                :value="field.replace(/\s+/g, '')"
              >
                {{ getTitleCase(field) }}
              </option>
            </select>
          </div>
          <div
            class="col-sm-12 col-md-12 col-lg-7"
            style="max-height: 500px; overflow: auto"
          >
            <div class="input-group">
              <input
                type="text"
                id="form1"
                placeholder="Search Employee"
                class="form-control"
                v-model="search"
              />
              <button
                type="button"
                @click="searchEmployee()"
                class="btn btn-primary"
              >
                <i class="fas fa-search"></i>
              </button>
            </div>
            <div
              class="row card p-2 emp-list"
              v-for="(emp, index) in employees"
              :key="index"
            >
              <div class="row" @click="viewEmp(emp)">
                <div class="col-sm-3">
                  <img
                    src="../assets/Images/avatar.png"
                    height="50"
                    width="50"
                    alt="profile"
                    class=""
                  />
                </div>

                <div class="col-sm-9">
                  <label class="text-dark"
                    ><b>{{
                      getTitleCase(emp.firstname + " " + emp.lastname)
                    }}</b></label
                  ><br />
                  <label class="">{{ emp.department }}</label>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Services from "../services/EmployeeServices.js";
export default {
  data() {
    return {
      fields: ["first name", "last name", "department", "email"],
      search: "",
      field: "",
      employees: [],
      emp: {},
    };
  },
  methods: {
    async searchEmployee() {
      let obj = {};
      obj[this.field] = this.search;
      console.log(obj);
      await Services.SearchEmployee(obj).then((data) => {
        console.log(data);
        if (data != null) {
          this.employees = data;
        } else {
          this.employees = [];
        }
      });
    },
    viewEmp(emp) {
      this.emp = emp;
      this.$refs["modal-view"].show();
    },
    clearSearch() {
      this.employees = [];
      this.search = "";
      this.field = "";
    },
  },
};
</script>
