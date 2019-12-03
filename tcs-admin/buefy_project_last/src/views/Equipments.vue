<template>
    <section class="p-3">
        <confirm ref="conf" @clicked="deleteEquip"></confirm>
        <b-table
                :data="isEmpty ? [] : items"
                :striped="isStriped"
                :narrowed="isNarrowed"
                :hoverable="isHoverable"
                :loading="isLoading"
                :mobile-cards="hasMobileCards">

            <template slot-scope="props">
                <b-table-column field="ID" label="ID" width="40" numeric>
                    {{ props.row.ID }}
                </b-table-column>

                <b-table-column field="Title" label="Title">
                    {{ props.row.Title }}
                </b-table-column>

                <b-table-column field="Type" label="Type">
                    {{ props.row.Type }}
                </b-table-column>

                <b-table-column field="SubType" label="SubType">
                    {{ props.row.SubType }}
                </b-table-column>

                <b-table-column field="BuyPrice" label="BuyPrice">
                    {{ props.row.BuyPrice }}
                </b-table-column>

                <b-table-column field="SellPrice" label="SellPrice">
                    {{ props.row.SellPrice }}
                </b-table-column>

                <b-table-column field="Reputation" label="Reputation">
                    {{ props.row.Reputation }}
                </b-table-column>

                <b-table-column field="Damage" label="Damage">
                    {{ props.row.Damage }}
                </b-table-column>

                <b-table-column field="Armor" label="Armor">
                    {{ props.row.Armor }}
                </b-table-column>

                <b-table-column field="Air" label="Air">
                    {{ props.row.Air }}
                </b-table-column>

                <b-table-column field="Mine" label="Mine">
                    {{ props.row.Mine }}
                </b-table-column>

                <b-table-column field="Time" label="Time">
                    {{ props.row.Time }}
                </b-table-column>

                <b-table-column field="SocketType" label="SocketType">
                    {{ props.row.SocketType }}
                </b-table-column>

                <b-table-column field="Sockets" label="Sockets">
                    {{ props.row.Sockets }}
                </b-table-column>

                <b-table-column field="" label="   ">
                    <b-button tag="router-link"
                              :to="{ path: `/equipments/update/${props.row.ID}` }"
                              type="is-link is-info"
                              size="is-small"
                              expanded
                    >
                        Update
                    </b-button>
                </b-table-column>

                <b-table-column field="" label="  ">
                    <b-button @click="confirmRemove(props.row.ID)" type="is-light" size="is-small" expanded>Remove</b-button>
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
        <div  class="pt-2">
            <b-button tag="router-link"
                      to="/equipments/add"
                      type="is-link is-success"
            >
                Add equip
            </b-button>
        </div>

    </section>
</template>

<script>
    import { mapGetters } from 'vuex';
    import Confirm from "../components/Confirm";

    export default {
        name: 'equipments',
        components: {
            'confirm': Confirm,
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
            }
        },
        created() {
            this.fetchData();
        },
        computed: {
            ...mapGetters("equips", [
                'items',
            ]),
            // ...mapGetters([
            //     'items',
            // ]),
        },
        methods: {
            fetchData(){
                this.$store.dispatch('equips/getEquips')
            },
            confirmRemove(id){
                this.$refs.conf.confirm(id)
            },
            deleteEquip(id){
                this.$store.dispatch('equips/deleteEquip', id)
            }
        }
    }
</script>

<style lang="scss">
    @import "../styles/_variables.scss";
    @import "../../node_modules/bulma/sass/utilities/_all.sass";

    table td:nth-child(1) {
        background-color: $violetpale;
    }
    .table thead th {
        background-color: $pink;
    }

    $table-row-hover-background-color: $hover;
    $narbar-hover-background-color: $hover;

    .b-table .table {
        background-color: $table;
    };
    /*// Import Bulma and Buefy styles*/
    $colors: (
    "light": ($pinkstrong, white),
    "primary": ($violet, white),
    "info": ($peachy, $white),
    "success": ($green, $success-invert),
    "warning": ($darkyellow, $white),
    "danger": ($pinkstrong, white),
    );
    .title{
        color: $violet !important;
    }
    .navbar{
        background-color: $darkgrey !important;
    }
    .navbar-item, .navbar-link{
        color: $green !important;
        font-size: 20px;
    }
    $navbar-item-hover-background-color:  $violetpale;

    .navbar-burger{
        color: $green !important;
    }
    $pagination-margin: 0.25rem !default;

    $pagination-color: $green !default;
    $pagination-border-color: $green !default;
    $pagination-active-color:  $peachy !default;
    $pagination-current-background-color: $violet !default;
    .navbar-menu{
        background-color: $darkgrey !important;
    }
    @import "../../node_modules/bulma";
    @import "../../node_modules/buefy/src/scss/buefy.scss";
</style>
