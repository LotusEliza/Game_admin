<template>
  <div id="app">
    <div id="navbar">
      <b-navbar toggleable="lg" type="dark" variant="info">
        <b-navbar-brand>ADMIN PANEL</b-navbar-brand>
        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
        <b-collapse id="nav-collapse" is-nav>
          <b-navbar-nav>
            <b-nav-item href="#"><router-link to="/">Home</router-link></b-nav-item>
            <b-nav-item href="#"><router-link to="/equipments">Equipments</router-link></b-nav-item>
            <b-nav-item href="#"><router-link to="/players">Players</router-link></b-nav-item>
          </b-navbar-nav>
          <!-- Right aligned nav items -->
          <b-navbar-nav class="ml-auto">
            <b-nav-form v-if="$route.meta.showNavigation">
                <b-form-input size="sm" v-model='id' class="mr-sm-2" placeholder="Search"></b-form-input>
                <b-button size="sm" @click="searchById" class="my-2 my-sm-0" type="submit">Search by ID</b-button>
            </b-nav-form>
            <b-nav-item-dropdown right>
              <!-- Using 'button-content' slot -->
              <template v-slot:button-content>
                <em>User</em>
              </template>
              <b-dropdown-item href="#">Profile</b-dropdown-item>
              <b-dropdown-item href="#">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown>
          </b-navbar-nav>
        </b-collapse>
      </b-navbar>
    </div>
    <router-view/>
  </div>
</template>

<script>
  export default {
    name: "norms",
    data() {
    return {
      id: '',
        }
      },
    methods: {
      searchById () {
        console.log(this.id)
        window.axios.get(`/player?id=${this.id}`).then((resp) => {
            // this.items = resp.data
            // console.log(resp.data);
            if(resp.data.Error){
              alert(resp.data.Error);
              this.id = ''
            }else{
              this.$router.push({ name: 'profile', params: resp.data })
              this.id = ''
            }
        }).catch(function (err) {
            console.log(err);
        });
      }
    },
  }
</script>

<style>
body{
  background-image: url("https://i.pinimg.com/originals/8f/ba/cb/8fbacbd464e996966eb9d4a6b7a9c21e.jpg")
}
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
#nav {
  padding: 30px;
}

#navbar a {

  color: #A4D8E2;
}

#navbar a.router-link-exact-active {
  font-weight: bold;
  color: #C6E7ED;
}
.table th{
  color: #42a3b8;
  font-weight: bold;
}
  /*.col{*/
  /*  color: red;*/
  /*  font-weight: bold;*/
  /*}*/
thead  {
  border-bottom: 4px solid   #a4a9ad;
}
table td:nth-child(1) {
  background-color: rgba(66, 163, 184, 0.4);
}


.table thead th {
  background-color: rgba(66, 163, 184, 0.2);
}

.table th:last-child, .table th:first-child {
  border-radius:6px;
}

.table td {
  background-color: rgba(255, 250, 250, 0.55);

}
.players table td:nth-last-child(-n+1)  {
  text-align:center;
}
</style>
