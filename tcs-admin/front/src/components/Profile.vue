<template>
    <div class="container">
        <div v-if="show">
            <div class="row">
                <div class="col-md-6 offset-md-3" style="text-align:center">
                    <h3>Player #{{player.Tg}}</h3>
                </div>
            </div>
            <div class="row">
                <div class="col-md-6 offset-md-3" style="text-align:center">
                    <div class="p-3" v-if="$mq === 'sm'">
                        <router-link
                                :to="{ name: 'norms', params: { tg: player.Tg } }"
                                tag="button" class="btn btn-danger mr-1" append>
                            Norms
                        </router-link>
                        <router-link
                                :to="{ name: 'players'}"
                                tag="button" class="btn btn-secondary">
                            Back
                        </router-link>
                    </div>
                    <table class="table table-borderless">
                        <tbody>

                        <!-- ********************************NAME********************************* -->
                        <tr>
                            <th scope="row">Name</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row">
                                        <div class="col">
                                            <input class="form-control form-control-sm"
                                                   v-model ="player.Name"
                                                   type="text"
                                                   placeholder=".form-control-sm">
                                        </div>
                                        <div class="col">
                                            <b-button @click="updatePlayer()" size="sm">Update</b-button>
                                        </div>
                                    </div>
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <th scope="row">Wallet</th>
                        </tr>

                        <!-- ********************************WALLET********************************* -->
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row buffer">
                                        <div class="col">
                                            Credits:
                                            <div class="def-number-input number-input safari_only">
                                                <input class="quantity" v-model.number="form.wallet.Credits" min="0" name="quantity" value="1" type="number">
                                            </div>
                                        </div>
                                        <div class="col">
                                            Gold:
                                            <div class="def-number-input number-input safari_only">
                                                <input class="quantity" v-model.number="form.wallet.Gold" min="0" name="quantity" value="1" type="number">
                                            </div>
                                        </div>
                                        <div class="col buffer2">
                                            <b-button @click="updateWallet()" size="sm">Update</b-button>
                                        </div>
                                    </div>
                                </div>
                            </td>
                        </tr>

                        <!-- ********************************INVENTORY********************************* -->
                        <tr>
                            <th scope="row">Inventory</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row buffer">
                                        <div class="col">
                                            <b-form-select v-model="selected" :options="firstInputOptions" size="sm" class="mt-3"></b-form-select>
                                        </div>
                                        <div class="col">
                                            <b-form-select v-model="secondInputSelected" :options="secondInputOptions" size="sm" id="mySelect" class="mt-3"></b-form-select>
                                        </div>
                                        <div class="col">
                                            <b-button align="center" @click="addItem" size="sm">Add</b-button>
                                        </div>
                                    </div>
                                    <ul style="list-style: none;">
                                        <li v-for="(item, key) in form.inventory">
                                            <div class="row buffer">
                                                <div class="col">
                                                    {{item.ItemType}}
                                                </div>
                                                <div class="col">
                                                    <div class="def-number-input number-input safari_only">
                                                        <input class="quantity" v-model="item.ItemValue" min="0" name="quantity" type="number">
                                                    </div>
                                                </div>
                                                <div class="col">
                                                    <b-button @click="removeItem(key)" size="sm">Remove</b-button>
                                                </div>
                                                <div class="col">
                                                    <b-button @click="updateItem(key)" size="sm">Update</b-button>
                                                </div>
                                            </div>
                                        </li>
                                    </ul>
                                </div>
                            </td>
                        </tr>

                        <!-- ********************************BALANCE********************************* -->
                        <tr>
                            <th scope="row">Balance</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="def-number-input number-input safari_only">
                                    <input class="quantity" v-model="form.balance" min="0" name="quantity" value="1" type="number">
                                </div>
                            </td>
                        </tr>

                        <!-- ********************************OXIGEN********************************* -->
                        <tr>
                            <th scope="row">Oxigen</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row">
                                        <div class="col">
                                            <span>Balloon:{{player.EquipB}}</span>
                                        </div>
                                        <div class="col">
                                            <div class="def-number-input number-input safari_only">
                                                <button @click="subOxigen"
                                                        onclick="this.parentNode.querySelector('input[type=number]').stepDown()"
                                                        class="minus"></button>
                                                <input class="quantity" v-model.number="form.oxigen" min="0" name="quantity" type="number">
                                                <button @click="plusOxigen"
                                                        onclick="this.parentNode.querySelector('input[type=number]').stepUp()"
                                                        class="plus"></button>
                                            </div>
                                        </div>
                                        <div class="col">
                                            <b-button size="sm" @click="fillUp">Fill up</b-button>
                                        </div>
                                        <div class="col">
                                            <b-button size="sm" @click="updateOxigen">Update</b-button>
                                        </div>
                                    </div>
                                </div>
                            </td>
                        </tr>

                        <!-- ********************************LOCATION********************************* -->
                        <tr>
                            <th scope="row ">Location</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row">
                                        <div class="col ">
                                            <span v-if="seen">X:{{player.PosX}}</span>
                                            <input v-if="!seen" class="form-control form-control-sm" v-model.number ="player.PosX" type="text" placeholder=".form-control-sm">
                                        </div>
                                        <div class="col ">
                                            <span v-if="seen">Y:{{player.PosY}}</span>
                                            <input v-if="!seen" class="form-control form-control-sm" v-model.number ="player.PosY" type="text" placeholder=".form-control-sm">
                                        </div>
                                        <div class="col ">
                                            <span v-if="seen">{{locations[player.Location]}}</span>
                                            <b-form-group id="input-group-3">
                                                <b-form-select
                                                        id="input-3"
                                                        v-if="!seen"
                                                        v-model="player.Location"
                                                        :options="options.locations"
                                                        required
                                                ></b-form-select>
                                            </b-form-group>
                                        </div>
                                        <div class="col">
                                            <b-button size="sm" @click="teleport" v-on:click="seen = !seen" v-if="seen">Teleport</b-button>
                                            <b-button size="sm" @click="updateLocation" v-on:click="seen = !seen" v-if="!seen">Update</b-button>
                                        </div>
                                    </div>
                                </div>
                            </td>
                        </tr>

                        <!-- ********************************CONSOLE********************************* -->
                        <tr>
                            <th scope="row">Console</th>
                        </tr>
                        <tr>
                            <td align="center">
                                <div class="container">
                                    <div class="row">
                                        <div class="col">
                                            <span>{{player.EquipC}}</span>
                                        </div>
                                    </div>
                                </div>
                            </td>
                        </tr>
                        <router-link v-on:click.native="saveTg(player.Tg)"
                                :to="{ name: 'norms', params: { tg: player.Tg } }"
                                tag="button" class="btn btn-danger mr-1" append>
                            Norms
                        </router-link>
                        <router-link
                                :to="{ name: 'players'}"
                                tag="button" class="btn btn-secondary">
                            Back
                        </router-link>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
        <div v-else>{{redPage()}}</div>
    </div>
</template>

<script>
    export default {
        props: ['item'],
        data() {
            return {
                show: true,
                seen: true,
                locations: ['Left Side', 'Right Side'],
                options: {
                    locations: [{ text: 'Left Side', value: 0 }, { text: 'Right Side', value: 1 },],
                    weapon: [
                        { value: 1, text: 'HGG' },
                        { value: 2, text: 'ТТ' },
                        { value: 3, text: 'МП-40'},
                    ],
                    armor: [
                        { value: 4, text: 'СКФ-98' },
                        { value: 5, text: 'СКФ-98а' },
                        { value: 6, text: 'СКФ-07'},
                    ],
                    metal: [
                        { value: 7, text: 'Железо' },
                        { value: 8, text: 'Медь' },
                        { value: 9, text: 'Титан'},
                    ],
                    gas: [
                        { value: 10, text: 'Метан' },
                        { value: 11, text: 'Водород' },
                        { value: 12, text: 'Гелий'},
                    ],
                    junk: [
                        { value: 13, text: 'Уголь' },
                        { value: 14, text: 'Песок' },
                        { value: 15, text: 'Вода'},
                        { value: 16, text: 'Камни'},
                        { value: 17, text: 'Кости'},
                        { value: 18, text: 'Шкура'},
                        { value: 19, text: 'Ветки'},
                    ],
                    device: [
                        { value: 20, text: 'Бур-1' },
                        { value: 21, text: 'Портативный NULL-T' },
                    ],
                },
                firstInputOptions: ['weapon', 'armor', 'metal', 'gas', 'junk', 'device'],
                selected: 'weapon',
                secondInputSelected: '',

                addInvent: '',
                player: this.item,
                form: {
                    selectedType: null,
                    selectedValue: null,
                    wallet: [],
                    inventory: '',
                    name: '',
                    oxigen: this.item.Air,
                    balance: this.item.Air
                },
            }
        },
        mounted() {
            this.getWallet()
            this.getInventory()
        },
        methods: {

            updateLocation(){
                console.log(this.player.Location + this.player.PosX + this.player.PosY)
                window.axios.post('/player/location/update', {
                    Tg: this.player.Tg,
                    Location: this.player.Location,
                    PosX: this.player.PosX,
                    PosY: this.player.PosY,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },

            updatePlayer(){
                console.log(this.player.Tg + this.player.Name)
                window.axios.post('/player/name', {
                    Tg: this.player.Tg,
                    Name: this.player.Name,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
                console.log(this.form.wallet.Gold)
            },

            updateWallet(){
                console.log("TG:" + this.player.Tg + this.form.wallet.Credits + this.form.wallet.Gold)
                window.axios.post('/player/wallet', {
                    Tg: this.player.Tg,
                    Credits: this.form.wallet.Credits,
                    Gold: this.form.wallet.Gold,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
                console.log(this.form.wallet.Gold)
            },

            addItem(key){
                console.log(this.secondInputSelected)
                window.axios.post('/player/inventory/add', {
                    tg: this.player.Tg,
                    itemtype: this.selected,
                    itemvalue: this.secondInputSelected,
                })
                    .then(function (response) {
                        console.log(response);

                    })
                    .catch(function (error) {
                        console.log(error);
                    });

                setTimeout( () => this.getInventory(), 100);
            },

            removeItem(key){
                console.log('removed'+key)
                window.axios.post('/player/inventory/remove', {
                    Tg: this.player.Tg,
                    ItemType: this.form.inventory[key].ItemType,
                    ItemValue: this.form.inventory[key].ItemValue,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
                setTimeout( () => this.getInventory(), 100);
            },

            updateItem(key){
                console.log('updated'+key+this.form.inventory[key].ItemType+this.form.inventory[key].ItemValue)
            },

            getWallet(){
                console.log('this is wallet get:'+this.player.Tg)
                window.axios.get(`/player/wallet?id=${this.player.Tg}`).then((resp) => {
                    console.log("this is axios getwallet:");
                    console.log(resp.data);
                    this.form.wallet = resp.data
                }).catch(function (err) {
                    console.log(err);
                });
            },

            getInventory(){
                window.axios.get(`/player/inventory?tg=${this.player.Tg}`).then((resp) => {
                    this.form.inventory = resp.data.Items
                }).catch(function (err) {
                    console.log(err);
                });
            },

            plusOxigen() {
                this.form.oxigen++
                console.log(this.form.oxigen)
            },

            subOxigen(){
                this.form.oxigen--
                console.log(this.form.oxigen)
            },

            redPage(){
                console.log('No data!')
                this.$router.push('players')
            },

            fillUp(){
                console.log("Filled up!")
                this.form.oxigen = 25

                window.axios.post('/player/air/update', {
                    Tg: this.player.Tg,
                    Air: this.form.oxigen,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },

            teleport(){
                console.log("Teleport!")
            },

            saveTg(Tg){
                // this.$store.dispatch('setPlayerId', Tg)
                localStorage.setItem('tg', Tg)
            },

            updateOxigen(){
                console.log("Updated!")
                window.axios.post('/player/air/update', {
                    Tg: this.player.Tg,
                    Air: this.form.oxigen,
                })
                    .then(function (response) {
                        console.log(response);
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            }

        },

        computed: {
            secondInputOptions(){
                return this.selected === 'weapon' ? this.options.weapon :
                    this.selected === 'armor' ? this.options.armor :
                        this.selected === 'device' ? this.options.device :
                            this.selected === 'metal' ? this.options.metal :
                                this.selected === 'gas' ? this.options.gas :
                                    this.selected === 'junk' ? this.options.junk :
                                        this.selected === 'device' ? this.options.device :
                                            null
            }
        },
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

    table td:nth-child(1) {
        background-color: rgba(255, 250, 250, 0.55)!important;
    }

    .table th:last-child, .table th:first-child{
        color: #42a3b8;
    }
    h3{
        background-color: rgba(108, 117, 125, 0.2);
        border-radius:6px;
        padding: 10px;
        margin-top: 10px;
        margin-bottom: 20px;
    }
</style>