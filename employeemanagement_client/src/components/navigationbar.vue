<template>
	<aside :class="`${is_expanded ? 'is-expanded' : ''}`">
		<div class="menu-toggle-wrap">
			<button class="menu-toggle" @click="ToggleMenu()">
				<span class="material-icons"><i class="fas fa-angle-double-right"></i></span>
			</button>
		</div>

		<h3>Menu</h3>
		<div class="menu">
			<router-link :to="link" class="button">
				<span class="material-icons"><i class="fa fa-home" aria-hidden="true"></i></span>
				<span class="text">Home</span>
			</router-link>
			<router-link to="/approvals" class="button">
				<span class="material-icons"><i class="fa-solid fa-bell"></i></span>
				<span class="text">Pending Approval</span>
			</router-link>
			<router-link to="/admin/add/employee" class="button" v-if="role=='admin' || role=='hr'">
				<span class="material-icons"><i class="fa-solid fa-user-plus"></i></span>
				<span class="text">Add Employee</span>
			</router-link>

            <router-link to="/admin/all/employees" class="button" v-if="role=='admin'">
				<span class="material-icons"><i class="fas fa-users"></i></span>
				<span class="text">Employee List</span>
			</router-link>
            <router-link to="/hr/all/employees" class="button" v-if="role=='hr'">
				<span class="material-icons"><i class="fas fa-users"></i></span>
				<span class="text">Employee List</span>
			</router-link>
			 <router-link to="/employee/profile" class="button" v-if="role=='hr' || role=='employee'">
				<span class="material-icons"><i class="fa-solid fa-user"></i></span>
				<span class="text">Profile</span>
			</router-link>
			 <router-link to="/apply/leaves" class="button" v-if="role=='hr' || role=='employee'">
				<span class="material-icons"><i class="fas fa-gifts"></i></span>
				<span class="text">Apply Leaves</span>
			</router-link>
		</div>

		<div class="flex"></div>
	</aside>
</template>

<script>
import { ref } from 'vue'
export default{
  data(){
    return{
      is_expanded:false,
      role:"",
      link:""
    }
  },
  created(){
     this.is_expanded = ref(localStorage.getItem("is_expanded") === true)
     this.role=localStorage.getItem('role')
     if(this.role=="admin"){
        this.link="/admin"
     }else if(this.role=="hr"){
         this.link="/hr"
     }else{
         this.link="/employee"
     }
  },
  methods:{
    ToggleMenu() {
	this.is_expanded = !this.is_expanded
	localStorage.setItem("is_expanded", this.is_expanded)
}
  }
  }
</script>

<style lang="scss" scoped>
aside {
	display: flex;
	flex-direction: column;
	background-color: var(--dark);
	color: var(--light);
	width: calc(2rem + 32px);
	overflow: hidden;
	min-height: 100vh;
	padding: 1rem;
	transition: 0.2s ease-in-out;
	.flex {
		flex: 1 1 0%;
	}
	.logo {
		margin-bottom: 1rem;
		img {
			width: 2rem;
		}
	}
	.menu-toggle-wrap {
		display: flex;
		justify-content: flex-end;
		position: relative;
		transition: 0.2s ease-in-out;
		.menu-toggle {
			transition: 0.2s ease-in-out;
			.material-icons {
				font-size: 1.3rem;
				color: var(--light);
				transition: 0.2s ease-out;
			}
			
			&:hover {
				.material-icons {
					color: var(--primary);
                   
					transform: translateX(0.5rem);
				}
			}
		}
	}
	h3, .button .text {
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
	}
	h3 {
		color: var(--grey);
		font-size: 0.875rem;
		margin-bottom: 0.5rem;
		text-transform: uppercase;
	}
	.menu {
		margin: 0 -1rem;
		.button {
			display: flex;
			align-items: center;
			text-decoration: none;
			transition: 0.2s ease-in-out;
			padding: 0.5rem 1rem;
			.material-icons {
				font-size: 1.3rem;
				color: var(--light);
				transition: 0.2s ease-in-out;
                margin-left: 10px;
			}
			.text {
				color: var(--light);
				transition: 0.2s ease-in-out;
			}
			&:hover {
				background-color: var(--dark-alt);
				.material-icons, .text {
					color: var(--primary);
				}
			}
			&.router-link-exact-active {
				background-color: var(--dark-alt);
				border-right: 5px solid var(--primary);
				.material-icons, .text {
					color: var(--primary);
				}
			}
		}
	}
	.footer {
		opacity: 0;
		transition: opacity 0.3s ease-in-out;
		p {
			font-size: 0.875rem;
			color: var(--grey);
		}
	}
	&.is-expanded {
		width: var(--sidebar-width);
		.menu-toggle-wrap {
			// top: -3rem;
			
			.menu-toggle {
				transform: rotate(-180deg);
			}
		}
		h3, .button .text {
			opacity: 1;
		}
		.button {
			.material-icons {
				margin-right: 1rem;
			}
		}
		.footer {
			opacity: 0;
		}
	}
	@media (max-width: 1024px) {
		position: absolute;
		z-index: 99;
	}
}
</style>