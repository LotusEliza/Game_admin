<template>
    <div class="p-3">
        <div class="tile is-ancestor">
            <div class="tile is-3 is-vertical is-parent">
            </div>
            <div class="tile is-6 is-vertical is-parent">
                <div class="tile is-child is-2">
                    <b-button @click="$router.go(-1)" type="is-info" >
                        Back
                    </b-button>
                </div>

<!-------------------------------------------------------------------------------------------------------->
<!--//********************************TABLE*******************************************************-->
<!-------------------------------------------------------------------------------------------------------->
                <b-table
                        :data="isEmpty ? [] : norms"
                        :striped="isStriped"
                        :narrowed="isNarrowed"
                        :hoverable="isHoverable"
                        :loading="isLoading"
                        :mobile-cards="hasMobileCards">

                    <template slot-scope="props">
                        <b-table-column field="Tg" label="Tg" width="40" numeric>
                            {{props.index+1}}
                        </b-table-column>

                        <b-table-column field="Resource" label="Resource" width="40" numeric centered>
                            <span  v-show="visible !== props.row.Date">
                                {{ options.resourceType[props.row.Resource-1].text }}
                            </span>
                            <b-select placeholder="Select a resource"
                                      :value="props.row.Resource"
                                      @input="updateResource"
                                      rounded
                                      size="is-small"
                                      v-show="visible === props.row.Date"
                            >
                                <option v-for="(option, key) in options.resourceType"
                                        :key="key"
                                        :value="option.value"
                                >
                                    {{option.text}}
                                </option>
                            </b-select>
                        </b-table-column>

                        <b-table-column field="Amount" label="Amount" centered>
                            <span  v-show="visible !== props.row.Date">{{ props.row.Amount }}</span>
                            <b-numberinput :value="props.row.Amount"
                                           @input="updateAmount"
                                           size="is-small"
                                           v-show="visible === props.row.Date"
                                           id="myDIV"
                            ></b-numberinput>
                        </b-table-column>

                        <b-table-column field="Date" label="Date">
                            {{ props.row.Date | formatDate }}
                        </b-table-column>

                        <b-table-column field="" label="   ">
                            <b-button @click="updateNorm(props.row.Date)"
                                      type="is-info"
                                      v-show="visible !== props.row.Date"
                                      :disabled="visible"
                                      size="is-small"
                                      expanded
                            >
                                Update
                            </b-button>
                            <b-button @click="saveNorm(props.row.Date)"
                                      type="is-info"
                                      v-show="visible === props.row.Date"
                                      size="is-small"
                                      expanded
                            >
                                Save
                            </b-button>
                        </b-table-column>

                        <b-table-column field="" label="  ">
                            <b-button @click="removeNorm(props.row.Date)"
                                      type="is-light"
                                      :disabled="visible"
                                      v-show="visible !== props.row.Date"
                                      size="is-small"
                                      expanded
                            >
                                Remove
                            </b-button>
                            <b-button @click="cancel(props.row, props.index)"
                                      type="is-light"
                                      v-show="visible === props.row.Date"
                                      size="is-small"
                                      expanded
                            >
                                Cancel
                            </b-button>
                        </b-table-column>
                    </template>
                    <template slot="empty">
                        <section class="section">
                            <div class="content has-text-grey has-text-centered">
                                <p>
                                    <b-icon
                                            icon="emoticon-sad"
                                            size="is-large">
                                    </b-icon>
                                </p>
                                <p>Nothing here.</p>
                            </div>
                        </section>
                    </template>
                </b-table>
<!-------------------------------------------------------------------------------------------------------->
<!--//********************************ADD NORM FORM*******************************************************-->
<!-------------------------------------------------------------------------------------------------------->
                <b-button @click="show = true"
                          type="is-primary"
                          expanded
                          v-show="!show"
                >
                    Add norm
                </b-button>
                <b-button @click="cancelAddNorm"
                          type="is-primary"
                          expanded
                          v-show="show"
                >
                    Cancel
                </b-button>
                <transition name="fade">
                    <div class="tile is-ancestor" v-if="show">
                        <div class="tile is-vertical is-parent">
                            <div class="tile is-child box">
                                <ValidationObserver ref="observer">
                                    <section slot-scope="{ validate }">
                                        <div class="columns is-mobile is-multiline">

                                            <div class="column is-narrow">
                                                <BSelectWithValidation rules="required"
                                                                       label="Resource"
                                                                       v-model="newNorm.Resource"
                                                >
                                                    <option v-for="(option, key) in options.resourceType"
                                                            :key="key"
                                                            :value="option.value"
                                                    >
                                                        {{option.text}}
                                                    </option>
                                                </BSelectWithValidation>
                                            </div>

                                            <div class="column is-narrow">
            <!--                                        <p class="subtitle">Date</p>-->
                                                <BInputWithValidation rules="required"
                                                                      type="text"
                                                                      label="Date"
                                                                      v-model="newNorm.Date"
                                                />
                                            </div>

                                            <div class="column">
                                                <label class="label">Amount</label>
                                                <b-numberinput v-model="newNorm.Amount"
                                                               rounded
                                                               size="is-small"
                                                ></b-numberinput>
                                            </div>
                                        </div>

                                        <div class="buttons pt-2">
                                            <b-button  @click="validate().then(saveNewNorm)" type="is-success">
                                                Add
                                            </b-button>
                                            <b-button @click="reset()" type="is-light">
                                                Reset
                                            </b-button>
                                        </div>
                                    </section>
                                </ValidationObserver>
                            </div>
                        </div>
                    </div>
                </transition>
            </div>
            <div class="tile is-3 is-vertical is-parent">
            </div>
        </div>
    </div>
</template>

<script>
    import { mapGetters } from 'vuex';
    import { ValidationObserver } from 'vee-validate';
    import BSelectWithValidation from '../components/inputs/BSelectWithValidation';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';

    export default {
        name: 'norms',
        components: {
            ValidationObserver,
            BSelectWithValidation,
            BInputWithValidation,
        },
        data() {
            return {
                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
                options: {
                    "resourceType": [
                        {text: 'gas', value: 1},
                        {text: 'metal', value: 2},
                        {text: 'oil', value: 3},
                        {text: 'send', value: 4},
                    ],
                },
                show: false,
                visible: false,
                newNorm: {
                    'Tg': Number(this.$route.params.tg),
                    'Resource': null,
                    'Date': this.callFunction(),
                    'Amount': null,
                },
                objAmount: null,
                amount: null,
                objResource: null,
            }
        },
        mounted() {
            this.getNorms();
        },
        computed: {
            ...mapGetters("norms", [
                'norms',
            ]),
        },

        methods:{
            getNorms(){
                this.$store.dispatch('norms/getNorms', this.$route.params.tg)
            },
            removeNorm(date){
                this.$store.dispatch('norms/deleteNorm', date);
            },
            updateNorm(date){
                this.visible = date;
                console.log('update norm!')
            },
            updateAmount (e) {
                this.amount = e;
                this.objAmount = {'Amount': e, 'Date': this.visible};
                this.$store.commit('norms/UPDATE_AMOUNT', this.objAmount)
            },
            updateResource (e) {
                this.objResource = {'Resource': e, 'Date': this.visible};
                this.$store.commit('norms/UPDATE_RESOURCE', this.objResource)
            },
            saveNorm(date){
                this.$store.dispatch('norms/updateNorm', date);
                this.visible = false;
            },
            cancel(norm, index){
                this.visible = false;
                let obj = {'index': index, 'norm': norm}
                this.$store.dispatch('norms/getNorm', obj)
            },
            cancelAddNorm(){
                this.show = false;
                this.$store.dispatch('norms/getNorm', this.$route.params.tg);
                this.reset();
            },
            reset(){
                this.newNorm.Resource = null;
                this.newNorm.Date = this.callFunction();
                this.newNorm.Amount = null;
                requestAnimationFrame(() => {
                    this.$refs.observer.reset();
                });
            },
            callFunction: function () {
                var currentDateWithFormat = new Date().toJSON().slice(0, 10).replace(/-/g, '-');
                return currentDateWithFormat
            },
            saveNewNorm(){
                let copiedObject = Object.assign({}, this.newNorm);
                this.$store.dispatch('norms/addNorm', copiedObject);
            }
        }
    }
</script>

<style>
    /*************Transition******************/
    .fade-enter-active, .fade-leave-active {
        transition: opacity .3s;
    }
    .fade-enter, .fade-leave-to /* .fade-leave-active до версии 2.1.8 */ {
        opacity: 0;
    }

</style>
