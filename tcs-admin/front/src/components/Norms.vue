<template>
    <div class="container">
        <div class="row">
            <div class="col-md-6 offset-md-3 header" style="text-align:center">
                <h3>Norms</h3>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6 offset-md-3" style="text-align:center">
                <table class="table table-borderless">
                    <tbody>
                    <tr>
                        <th scope="row">№</th>
                        <th scope="row">Resource</th>
                        <th scope="row">Amount</th>
                        <th scope="row">Date</th>
                        <th scope="row"></th>
                        <th scope="row"></th>
                    </tr>
                    <tr v-for="(item, key) in norms" :key="key">
                        <td scope="row">
                            {{key+1}}
                        </td>

    <!-- ************************** RESOURCE ************************** -->
                        <td scope="row">
                            <span  v-show="visible !== key">{{ options[item.Resource-1].name }}</span>
                            <b-form-select
                                    v-show="visible === key"
                                    v-model="item.Resource"
                                    :options="options"
                                    class="mb-3"
                                    value-field="item"
                                    text-field="name"
                                    disabled-field="notEnabled"
                            ></b-form-select>
                        </td>

    <!-- ************************** AMOUNT ************************** -->
                        <td scope="row">
                            <span  v-show="visible !== key">{{ item.Amount }}</span>
                            <b-form-input v-model="item.Amount" v-show="visible === key"></b-form-input>
                        </td>

    <!-- ************************** DATE ************************** -->
                        <td scope="row">
                            <span>{{ item.Date | formatDate }}</span>
                        </td>

    <!-- ************************** BUTTONS (REMOVE, UPDATE) ************************** -->
                        <td scope="row">
                            <button v-on:click="removeNorm(item, key)"
                                    class="btn btn-secondary btn-sm"
                                    :disabled="visible"
                                    v-show="visible !== key">
                                Remove
                            </button>
                            <button v-on:click="cancelUpdate(item, key)"
                                    class="btn btn-secondary btn-sm"
                                    v-show="visible === key">
                                Cancel
                            </button>
                        </td>
                        <td scope="row">
                            <button type="button"
                                    @click="edit(item, key)"
                                    class="btn btn-secondary btn-sm"
                                    v-show="visible !== key"
                                    :disabled="visible"
                            >Update</button>
                            <button type="button"
                                    @click="updateNorm(item, key)"
                                    class="btn btn-secondary btn-sm"
                                    v-show="visible === key"
                            >Save</button>
                        </td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>
        <div class="row">
            <div class="col-md-6 offset-md-3" style="text-align:center">

    <!-- ************************** ADD NEW NORM BTN ************************** -->
                    <b-row class="p-3" >
                        <b-button class="btn btn-secondary btn-sm btn-block"
                                  @click="show=!show , color = true" v-if="!show">
                            Add new norm
                        </b-button>
                        <b-button class="btn btn-secondary btn-sm btn-block"
                                  @click="cancel()" v-if="show" >
                            Cancel
                        </b-button>
                    </b-row>
            </div>
        </div>

        <div class="row">
            <div class="col-md-6 offset-md-3" style="text-align:center">

                <div  class="container"
                      :class="{'color-form': color}">
                    <transition name="fade">
                    <b-form @submit="addNorm" @reset="onReset" v-if="show"  class="stylelabel">

                        <div  >
                            <div class="row buffer">

    <!--**************************** RESOURCE ********************************-->
                                <div class="col">
                                    <b-form-group
                                            id="input-group-2"
                                            label="Resourse:"
                                            label-for="input-2"
                                    >
                                    <b-form-select
                                            class="mb-3"
                                            value-field="value"
                                            id="input-2"
                                            v-model="form.Resource"
                                            :options="optionsAdd"
                                    ></b-form-select>
                                    </b-form-group>
                                </div>

    <!-- ************************** DATE ************************** -->
                                <div class="col">
                                    <b-form-group
                                            id="input-group-1"
                                            label="Date:"
                                            label-for="input-1"
                                    >
                                    <b-form-input
                                            id="input-1"
                                            v-model="form.Date"
                                    ></b-form-input>
                                    </b-form-group>
                                </div>

    <!--**************************** AMOUNT ********************************-->
                                <div class="col">
                                    <b-form-group
                                            id="input-group-3"
                                            label="Amount:"
                                            label-for="input-3"
                                    >
                                    <b-form-input
                                            id="input-3"
                                            v-model.number="form.Amount"
                                            type="number"
                                            placeholder="Enter amount"
                                        ></b-form-input>
                                    </b-form-group>
                                </div>
                            </div>
                        </div>
                        <div class="buttons p-3">
                            <b-button type="submit" variant="btn btn-secondary mr-1">Add</b-button>
                            <b-button type="reset" variant="danger mr-1">Reset</b-button>
                        </div>
                    </b-form>
                    </transition>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import { mapState } from 'vuex';
    import { mapGetters } from 'vuex';
export default {
    data() {
        return {
            colorForm: 'color-form',
            color: false,
            show: false,
            visible: false,
            options: [
                { item: 1, name: 'metal' },
                { item: 2, name: 'water' },
                { item: 3, name: 'gas' },
                { item: 4, name: 'oil' },
                { item: 5, name: 'sand' },
            ],
            optionsAdd: [
                { value: 1, text: 'metal' },
                { value: 2, text: 'water' },
                { value: 3, text: 'gas' },
                { value: 4, text: 'oil' },
                { value: 5, text: 'sand' },
            ],
            form: {
                Tg: null,
                Resource: null,
                Amount: null,
                Date:  this.callFunction(),
            },
        }
    },
    mounted() {
        if(!this.$route.params.tg){
            this.$store.dispatch('setPlayerId', localStorage.getItem('tg'))
        }else {
            this.$store.dispatch('setPlayerId', this.$route.params.tg)
        }
        this.$store.dispatch('getNorms', this.$store.state.norms.playerId)
    },
    computed:{
        ...mapGetters([
            'norms',
            'tg',
        ]),
        // ...mapState({
        //     norms: state => state.norms.norms,
        //     tg: state => state.norms.playerId
        // }),
        // norms(){return this.$store.state.norms.norms},
    },
    methods: {
        removeNorm(norm, key) {
            this.$store.dispatch('deleteNorm', norm)
        },
        edit(item, key) {
            this.visible = key;
        },
        updateNorm(norm, key) {
            this.$store.dispatch('updateNorm', norm);
            this.visible = false
        },
        addNorm() {
            this.form.Tg = Number(this.$store.state.norms.playerId);
            this.$store.dispatch('addNorm', this.form);
            this.$store.dispatch('getNorms', this.$store.state.norms.playerId);
        },
        callFunction: function () {
            var currentDateWithFormat = new Date().toJSON().slice(0, 10).replace(/-/g, '-');
            return currentDateWithFormat
        },
        cancel() {
            this.color = false;
            this.show = !this.show;
            this.form.resource = null;
            this.form.amount = null;
            this.form.date = this.callFunction();
        },
        cancelUpdate() {
            this.$store.dispatch('getNorms', this.$store.state.playerId);
            this.visible = false
        },
        onReset(evt) {
            evt.preventDefault();
            // Reset our form values
            this.form.Resource = '';
            this.form.Amount = '';
            this.form.Date = this.callFunction();
        },
    }
}
</script>
<style scoped>
    .color-form {
        background-color: rgba(255, 250, 250, 0.55);
        padding: 10px;
    }
    .table > tbody > tr:first-child {
        background-color: rgba(66, 163, 184, 0.2);
        border-bottom: 4px solid   #a4a9ad;
    }
    h3{
        background-color: rgba(108, 117, 125, 0.2);
        border-radius:6px;
        padding: 10px;
        margin-top: 10px;
        margin-bottom: 20px;
    }
   /*************Transition******************/
    .fade-enter-active, .fade-leave-active {
        transition: opacity .5s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
        opacity: 0;
    }

</style>