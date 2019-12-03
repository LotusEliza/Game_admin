<template>
  <div class="container">
    <div v-if="item.ID">
      <div class="row">
        <div class="col-md-6 offset-md-3 header p-2" style="text-align:center">
            <h2>Update equipment</h2>
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
                    v-model="item.Title"
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
                        v-model="item.Type"
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
                      v-model="item.SubType"
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
                          v-model.number="item.BuyPrice"
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
                          v-model.number="item.SellPrice"
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
                        v-model.number="item.Reputation"
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
                        v-model.number="item.Damage"
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
                        v-model.number="item.Armor"
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
                        v-model.number="item.Air"
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
                        v-model.number="item.Mine"
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
                        v-model.number="item.Time"
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
                      v-model="item.SocketType"
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
                    v-model.number="item.Sockets"
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
    <div v-else>{{redPage()}}</div>
  </div>
</template>
<script>
  export default {
      props: ['exampleProp'],
    data() {
      return {
        item: this.exampleProp,
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
      mounted(){

      },
    methods: {
      onSubmit(evt) {
        evt.preventDefault()
          this.$store.dispatch('updateEquip', this.item)
      },
      onReset(evt) {
        evt.preventDefault()
        // Reset our form values
        this.item.Type = ''
        this.item.SubType = ''
        this.item.BuyPrice = null
        this.item.SellPrice = null
        this.item.Reputation = null
        this.item.Damage = null
        this.item.Armor = null
        this.item.Air = null
        this.item.Mine = null
        this.item.Time = null
        this.item.SocketType = null
        this.item.Sockets = null
        // Trick to reset/clear native browser form validation state
        this.show = false
        this.$nextTick(() => {
          this.show = true
        })
      },
    redPage(){
      console.log('No data!')
      this.$router.push('equipments') 
    },
    }
  }
</script>
<style scoped>

.number-input input[type="number"] {
  -webkit-appearance: textfield;
  -moz-appearance: textfield;
  appearance: textfield;
}

.number-input input[type=number]::-webkit-inner-spin-button,
.number-input input[type=number]::-webkit-outer-spin-button {
  -webkit-appearance: none;
}

.number-input button {
  -webkit-appearance: none;
  background-color: transparent;
  border: none;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  margin: 0;
  position: relative;
}

.number-input button:before,
.number-input button:after {
  display: inline-block;
  position: absolute;
  content: '';
  height: 2px;
  transform: translate(-50%, -50%);
}

.number-input button.plus:after {
  transform: translate(-50%, -50%) rotate(90deg);
}

.number-input input[type=number] {
  text-align: center;
}

.number-input.number-input {
  border: 1px solid #ced4da;
  width: 10rem;
  border-radius: .25rem;
}

.number-input.number-input button {
  width: 2.6rem;
  height: .7rem;
}

.number-input.number-input button.minus {
  padding-left: 10px;
}

.number-input.number-input button:before,
.number-input.number-input button:after {
  width: .7rem;
  background-color: #495057;
}

.number-input.number-input input[type=number] {
  max-width: 4rem;
  padding: .5rem;
  border: 1px solid #ced4da;
  border-width: 0 1px;
  font-size: 1rem;
  height: 2rem;
  color: #495057;
}

@media not all and (min-resolution:.001dpcm) {
  @supports (-webkit-appearance: none) and (stroke-color:transparent) {

    .number-input.def-number-input.safari_only button:before,
    .number-input.def-number-input.safari_only button:after {
      margin-top: -.3rem;
    }
  }
}

th{
  background-color: rgba(66, 163, 184, 0.2);
  border-radius:6px;
}

.def-number-input number-input safari_only{
  text-align: center;
}

.table th, .table td {
    padding: 0.5rem;
}

.container {
  padding: 10px;
}

.buffer {
  padding: 7px;
}

.buffer2 {
  padding-top: 22px;
}

.mt-3 {
    margin-top: 0rem !important;
}

.col-md-6{
     background-color: rgba(255, 250, 250, 0.55);
     border-radius:6px;
     /* padding: 5px; */
}

.header{
   background-color: rgba(66, 163, 184, 0.2);
}

.stylelabel, h2 {
    color: rgb(66, 163, 184);
}

</style>