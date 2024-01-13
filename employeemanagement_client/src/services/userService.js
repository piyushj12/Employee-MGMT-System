import axios from "axios";
// import router from '../router'
// import errorjs from "../Javascript/error"
const HOST = process.env.HOST || "0.0.0.0";
const PORT = process.env.PORT || "4700";

const apiClient = axios.create({
  baseURL: `http://${HOST}:${PORT}`,
  withCredentials: false,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
     Authorization:localStorage.getItem("token"),
  },
 
});

export default {
    async loginService(payload) {
        try{
          let res = await apiClient.post("/o/login",payload);
          return res.data;
        }catch(error){
          return error.response.data.message;
      }
    },
    async addEmployeeService(payload) {
      try{
        let res = await apiClient.post("/ah/add",payload);
        return res.data;
      }catch(error){
        return error.response.data.message;
    }
  },
  async getAllEmployeesService(payload) {
    try{
      let res = await apiClient.post("/ah/get/all/employees",payload);
      return res.data;
    }catch(error){
      return error.response.data.message;
  }
},
async getManagerService() {
  try{
    let res = await apiClient.get("/ah/get/managers")
    return res.data;
  }catch(error){
    return error.response.data;
  }
},
async AddField(payload) {
  try{
    let res = await apiClient.put("/a/update/array",payload)
    return res.data;
  }catch(error){
    return error.response.data;
  }
},
async getCompanyData() {
  try{
    let res = await apiClient.get("/r/get/company/data")
    return res.data;
  }catch(error){
    return error.response.data;
  }
},
async ResetData() {
  try{
    let res = await apiClient.put("/a/reset/employee/data")
    return res.data;
  }catch(error){
    return error.response.data;
  }
}
};
