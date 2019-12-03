<template>
  <div class="container">
      <div class="row">
        <div class="col-md-6 offset-md-3 header p-2" style="text-align:center">
            <h2>Add new equipment</h2>
        </div>
        <div class="col-md-6 offset-md-3 p-4" style="text-align:left">
            <b-form @submit="onSubmit" @reset="onReset" v-if="show"  class="stylelabel">

<!--**************************** TITLE ********************************-->
                <b-form-group
                    id="input-group-1"
                    label="Title:"
                    label-for="input-1"
                >
                    <b-form-input
                    id="input-1"
                    v-model="form.title"
                    required
                    placeholder="Enter title"
                    ></b-form-input>
                </b-form-group>
                <div class="container">
                  <div class="row buffer">

<!--**************************** TYPE ********************************-->
                    <div class="col">
                      <b-form-group 
                        id="input-group-13" 
                        label="Type:" 
                        label-for="input-13"
                      >
                      <b-form-select
                        id="input-13"
                        v-model="form.type"
                        :options="options.type"
                        required
                      ></b-form-select>
                      </b-form-group>
                    </div>

<!--**************************** SUB TYPE ********************************-->
                    <div class="col">
                    <b-form-group 
                      id="input-group-13" 
                      label="Sub Type:" 
                      label-for="input-13"
                    >
                    <b-form-select
                      id="input-13"
                      v-model="form.subType"
                      :options="options.subType"
                      required
                    ></b-form-select>
                    </b-form-group>
                    </div>
                  </div>
                  <div class="row buffer">

<!--**************************** BUY PRICE SELL PRICE ********************************-->
                    <div class="col">
                      <b-form-group
                          id="input-group-4"
                          label="Buy Price:"
                          label-for="input-4"
                      >
                          <b-form-input
                          id="input-4"
                          v-model.number="form.buyPrice"
                          required
                          type="number"
                          placeholder="Enter buy price"
                          ></b-form-input>
                      </b-form-group>
                  </div>

<!--**************************** SELL PRICE ********************************-->
                  <div class="col">
                      <b-form-group
                          id="input-group-5"
                          label="Sell Price:"
                          label-for="input-5"
                      >
                          <b-form-input
                          id="input-5"
                          v-model.number="form.sellPrice"
                          required
                          type="number"
                          placeholder="Enter sell price"
                          ></b-form-input>
                      </b-form-group>
                  </div>
                </div>
                <div class="row buffer">

<!--**************************** REPUTATION ********************************-->
                  <div class="col">
                    <b-form-group
                        id="input-group-6"
                        label="Reputation:"
                        label-for="input-6"
                    >
                    <b-form-input
                        id="input-6"
                        v-model.number="form.reputation"
                        required
                        type="number"
                        placeholder="Enter reputation"
                        ></b-form-input>
                    </b-form-group>
                  </div>

<!--**************************** DAMAGE ********************************-->
                  <div class="col">
                    <b-form-group
                        id="input-group-7"
                        label="Damage:"
                        label-for="input-7"
                    >
                        <b-form-input
                        id="input-7"
                        v-model.number="form.damage"
                        required
                        type="number"
                        placeholder="Enter damage"
                        ></b-form-input>
                    </b-form-group>
                  </div>

<!--**************************** ARMOR ********************************-->
                  <div class="col">
                    <b-form-group
                        id="input-group-8"
                        label="Armor:"
                        label-for="input-8"
                    >
                        <b-form-input
                        id="input-8"
                        v-model.number="form.armor"
                        required
                        type="number"
                        placeholder="Enter armor"
                        ></b-form-input>
                    </b-form-group>
                  </div>
                </div>
                <div class="row buffer">

<!--**************************** AIR ********************************-->
                  <div class="col">
                    <b-form-group
                        id="input-group-9"
                        label="Air:"
                        label-for="input-9"
                    >
                        <b-form-input
                        id="input-9"
                        v-model.number="form.air"
                        required
                        type="number"
                        placeholder="Enter air"
                        ></b-form-input>
                    </b-form-group>
                  </div>

<!--**************************** MINE ********************************-->
                  <div class="col">
                    <b-form-group
                        id="input-group-10"
                        label="Mine:"
                        label-for="input-10"
                    >
                        <b-form-input
                        id="input-10"
                        v-model.number="form.mine"
                        required
                        type="number"
                        placeholder="Enter mine"
                        ></b-form-input>
                    </b-form-group>
                  </div>

<!--**************************** TIME ********************************--> 
                  <div class="col">
                    <b-form-group
                        id="input-group-11"
                        label="Time:"
                        label-for="input-11"
                    >
                        <b-form-input
                        id="input-11"
                        v-model.number="form.time"
                        required
                        type="number"
                        placeholder="Enter time"
                        ></b-form-input>
                    </b-form-group>
                  </div>
                </div>
                <div class="row buffer">

<!--**************************** SOCKET TYPE ********************************-->
                    <div class="col">
                    <b-form-group 
                      id="input-group-13" 
                      label="Socket Type:" 
                      label-for="input-13"
                    >
                    <b-form-select
                      id="input-13"
                      v-model="form.socketType"
                      :options="options.socketType"
                      required
                    ></b-form-select>
                    </b-form-group>
                    </div>

<!--**************************** SOCKETS ********************************--> 
                    <div class="col">
                    <b-form-group
                        id="input-group-12"
                        label="Sockets:"
                        label-for="input-12"
                    >
                    <b-form-input
                    id="input-12"
                    v-model.number="form.sockets"
                    required
                    type="number"
                    placeholder="Enter sockets"
                    ></b-form-input>
                   </b-form-group>
                    </div>
                 </div>
              </div>
              <b-button type="submit" variant="btn btn-secondary mr-1">Submit</b-button>
              <b-button type="reset" variant="danger mr-1">Reset</b-button>
              <router-link  
                            :to="{ name: 'equipments'}"
                            tag="button" class="btn btn-secondary mr-1">
                                Back
              </router-link>
            </b-form>
        </div>
      </div>
    </div>
</template>

<script>
  export default {
    data() {
      return {
        form: {
          title: '',
          type: '',
          subType: null,
          buyPrice: null,
          sellPrice: null,
          reputation: null,
          damage: null,
          armor: null,
          air: null,
          mine: null,
          time: null,
          socketType: null,
          sockets: null,
        },
        options: {
            "type": ['weapon', 'armor', 'balloon', 'console'],
            "subType": ['boer', 'nullt'],
            "socketType": [
                {text: 'damage', value: 1}, 
                {text: 'armor', value: 1},
                {text: 'hp', value: 2}, 
                {text: 'mine', value: 3}],
        },
        show: true
      }
    },
    methods: {
      onSubmit(evt) 
      {
        evt.preventDefault()
          this.$store.dispatch('addEquip', this.form)
          // this.$router.push({ path: '/equipments'});

        // alert(JSON.stringify(this.form))
      // window.axios.post('/equipments/add/', {
      //   Title: this.form.title,
      //   Type: this.form.type,
      //   SubType: this.form.subType,
      //   BuyPrice: this.form.buyPrice,
      //   SellPrice: this.form.sellPrice,
      //   Reputation: this.form.reputation,
      //   Damage: this.form.damage,
      //   Armor: this.form.armor,
      //   Air: this.form.air,
      //   Mine: this.form.mine,
      //   Time: this.form.time,
      //   SocketType: this.form.socketType,
      //   Sockets: this.form.sockets,
      // })
      // .then(function (response) {
      //
      //   console.log(response);
      // })
      // .catch(function (error) {
      //   console.log(error);
      // });

      },
      onReset(evt) {
        evt.preventDefault()
        // Reset our form values
        this.form.title = ''
        this.form.type = ''
        this.form.subType = ''
        this.form.buyPrice = null
        this.form.sellPrice = null
        this.form.reputation = null
        this.form.damage = null
        this.form.armor = null
        this.form.air = null
        this.form.mine = null
        this.form.time = null
        this.form.socketType = null
        this.form.sockets = null
        // Trick to reset/clear native browser form validation state
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      }
    }
  }
</script>
<style scoped>

.container {
  padding: 10px;
}

.col-md-6{
     background-color: rgba(255, 250, 250, 0.55);
     border-radius:6px;
}

.header {
   background-color: rgba(66, 163, 184, 0.2);
}

.stylelabel, h2 {
    color: rgb(66, 163, 184);
}
</style>