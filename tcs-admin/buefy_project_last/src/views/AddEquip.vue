<template>
    <div class="p-3">
        {{form.Title}}
        <ValidationObserver ref="observer">
            <section class="section" slot-scope="{ validate }">
                <div class="tile is-ancestor">
                    <div class="tile is-2 is-vertical is-parent">
                    </div>
                    <div class="tile is-vertical is-parent">
                        <div class="tile is-child box">
                            <p class="title">Add Equipment</p>
                            <section>
                                <div class="columns is-mobile is-multiline">
                                    <div class="column is-narrow">
                                        <BInputWithValidation rules="required"
                                                              type="text"
                                                              label="Title"
                                                              vid="password"
                                                              v-model="form.Title"
                                                              placeholder="Enter title"
                                        />
                                    </div>
                                    <div class="column is-narrow">
                                        <BSelectWithValidation rules="required"
                                                               label="Type"
                                                               v-model="form.Type"
                                        >
                                            <option v-for="(option, key) in options.type" :key="key">
                                                {{option}}
                                            </option>
                                        </BSelectWithValidation>
                                    </div>
                                    <div class="column is-narrow">
                                        <BSelectWithValidation rules="required"
                                                               label="Sub Type"
                                                               v-model="form.SubType"
                                                               placeholder="Select a sub type"
                                        >
                                            <option v-for="(option, key) in options.subType" :key="key">
                                                {{option}}
                                            </option>
                                        </BSelectWithValidation>
                                    </div>
                                    <div class="column is-narrow">
                                        <BSelectWithValidation rules="required" label="Socket Type" v-model="form.SocketType">
                                            <option v-for="(option, key) in options.socketType" :key="key" :value="option.value">
                                                {{option.text}}
                                            </option>
                                        </BSelectWithValidation>
                                    </div>
                                </div>
                                <div class="columns is-mobile is-multiline">
                                    <div class="column">
                                        <b-field label="Sockets">
                                            <b-numberinput v-model="form.Sockets"
                                                           placeholder="Enter sockets"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Buy price">
                                            <b-numberinput v-model="form.BuyPrice"
                                                           placeholder="Enter buy price"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Sell price">
                                            <b-numberinput v-model="form.SellPrice"
                                                           placeholder="Enter sell price"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                </div>
                                <div class="columns is-mobile is-multiline">
                                    <div class="column">
                                        <b-field label="Reputation">
                                            <b-numberinput v-model="form.Reputation"
                                                           placeholder="Enter reputation"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Damage">
                                            <b-numberinput v-model="form.Damage"
                                                           placeholder="Enter damage"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Armor">
                                            <b-numberinput v-model="form.Armor"
                                                           placeholder="Enter armor"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                </div>
                                <div class="columns is-mobile is-multiline">
                                    <div class="column">
                                        <b-field label="Air">
                                            <b-numberinput v-model="form.Air"
                                                           placeholder="Enter air"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Mine">
                                            <b-numberinput v-model="form.Mine"
                                                           placeholder="Enter mine"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                    <div class="column">
                                        <b-field label="Time">
                                            <b-numberinput v-model="form.Time"
                                                           placeholder="Enter time"
                                                           rounded
                                                           size="is-small"
                                            ></b-numberinput>
                                        </b-field>
                                    </div>
                                </div>
                            </section>
                            <section>
                                <div class="buttons pt-2">
                                    <b-button  @click="validate().then(submit)" type="is-success">Submit</b-button>
                                    <b-button  @click="onReset()" type="is-info">Reset</b-button>
                                    <b-button tag="router-link"
                                              :to="{ path: `/equipments` }"
                                              type="is-link is-light"
                                    >
                                        Back
                                    </b-button>
                                </div>
                            </section>
                        </div>
                    </div>
                    <div class="tile is-2 is-vertical is-parent">
                    </div>
                </div>
            </section>
        </ValidationObserver>
    </div>
</template>

<script>
    import { ValidationObserver } from 'vee-validate';
    import BSelectWithValidation from '../components/inputs/BSelectWithValidation';
    import BInputWithValidation from '../components/inputs/BInputWithValidation';

    export default {
        name: 'addEquip',
        components: {
            ValidationObserver,
            BSelectWithValidation,
            BInputWithValidation,
        },
        data() {
            return {
                form: {
                    Title: '',
                    Type: null,
                    SubType: null,
                    BuyPrice: null,
                    SellPrice: null,
                    Reputation: null,
                    Damage: null,
                    Armor: null,
                    Air: null,
                    Mine: null,
                    Time: null,
                    SocketType: null,
                    Sockets: null,
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
            }
        },
        methods: {
            onReset() {
                // Reset our form values
                this.form.Title = '';
                this.form.Type = '';
                this.form.SubType = '';
                this.form.BuyPrice = null;
                this.form.SellPrice = null;
                this.form.Reputation = null;
                this.form.Damage = null;
                this.form.Armor = null;
                this.form.Air = null;
                this.form.Mine = null;
                this.form.Time = null;
                this.form.SocketType = null;
                this.form.Sockets = null;
                requestAnimationFrame(() => {
                    this.$refs.observer.reset();
                });
            },
            async submit () {
                console.log(this.form);
                try {
                    await   this.$store.dispatch('equips/addEquip', this.form);
                } catch (error) {
                    console.log('error')
                } finally {
                    await this.$router.push('/equipments');
                }
                // this.$store.dispatch('addEquip', this.form);
                // this.$router.push({path: '/equipments'});
            },
        }
    }
</script>

<style>
</style>
