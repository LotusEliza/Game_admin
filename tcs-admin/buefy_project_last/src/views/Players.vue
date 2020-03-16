<template>
    <section class="pd-3">
        <app-search class="is-hidden-desktop p-3"></app-search>

<!--****************************PAGINATION*******************************************-->
        <section>
            <b-pagination
                    :total="total_players"
                    :current.sync="current"
                    :range-before="rangeBefore"
                    :range-after="rangeAfter"
                    :order="order"
                    :size="size"
                    :rounded="isRounded"
                    :per-page="perPage"
                    :icon-prev="prevIcon"
                    :icon-next="nextIcon"
                    aria-next-label="Next page"
                    aria-previous-label="Previous page"
                    aria-page-label="Page"
                    aria-current-label="Current page">
            </b-pagination>
        </section>

<!--****************************TABLE*******************************************-->
        <b-table
                :data="isEmpty ? [] : players"
                :striped="isStriped"
                :narrowed="isNarrowed"
                :hoverable="isHoverable"
                :loading="isLoading"
                :mobile-cards="hasMobileCards"

        >

            <template slot-scope="props">
                <b-table-column field="Tg" label="Tg" width="40" numeric>
                    {{ props.row.Tg }}
                </b-table-column>

                <b-table-column field="Name" label="Name">
                    {{ props.row.Name }}
                </b-table-column>

                <b-table-column field="Story" label="Story">
                    {{ props.row.Story }}
                </b-table-column>

                <b-table-column field="Faction" label="Faction">
                    {{ props.row.Faction }}
                </b-table-column>

                <b-table-column field="Referrer" label="Referrer">
                    {{ props.row.Referrer }}
                </b-table-column>

                <b-table-column field="Location" label="Location">
                    {{ locations[props.row.Location] }}
                </b-table-column>

                <b-table-column field="PosX" label="PosX">
                    {{ props.row.PosX }}
                </b-table-column>

                <b-table-column field="PosY" label="PosY">
                    {{ props.row.PosY }}
                </b-table-column>

                <b-table-column field="HP" label="HP">
                    {{ props.row.HP }}
                </b-table-column>

                <b-table-column field="EquipW" label="EquipW">
                    {{ props.row.EquipW }}
                </b-table-column>

                <b-table-column field="EquipA" label="EquipA">
                    {{ props.row.EquipA }}
                </b-table-column>

                <b-table-column field="EquipB" label="EquipB">
                    {{ props.row.EquipB }}
                </b-table-column>

                <b-table-column field="EquipC" label="EquipC">
                    {{ props.row.EquipC }}
                </b-table-column>

                <b-table-column field="Air" label="Air">
                    {{ props.row.Air }}
                </b-table-column>

                <b-table-column field="Registered" label="Registered">
                    {{ props.row.Registered }}
                </b-table-column>

                <b-table-column field="LastActive" label="LastActive">
                    {{ props.row.LastActive }}
                </b-table-column>

                <b-table-column field="" label="" >
                    <b-button tag="router-link"
                              :to="{ path: `/players/profile/${props.row.Tg}` }"
                              type="is-link is-primary"
                              class="is-small is-hidden-tablet"
                              expanded
                    >
                        Profile
                    </b-button>
                    <b-button tag="router-link"
                              :to="{ path: `/players/profile/${props.row.Tg}` }"
                              type="is-link is-primary"
                              class="is-small is-hidden-mobile"
                    >
                        Profile
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
    </section>
</template>

<script>
    import { mapGetters } from 'vuex';
    import Search from '../components/SearchPlayer.vue'
    import { locations } from '../utils/variables.js'

    export default {
        name: 'Players',
        components: {
            'app-search': Search,
        },
        data() {
            return {
            //************table******************
                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
            //*************pagination**************
                total: 10,
                current: 1,
                perPage: 10,
                rangeBefore: 3,
                rangeAfter: 1,
                order: 'is-centered',
                size: '',
                isRounded: true,
                prevIcon: 'arrow-left',
                nextIcon: 'arrow-right',

                locations: '',
            }
        },
        mounted(){
            this.locations = locations;
            this.$store.dispatch('players/getPlayersByPage', this.current);
        },
        computed:{
            ...mapGetters("players", [
                'total_players',
                'players',
            ]),
        },
        watch: {
            current: {
                handler: function() {
                    this.$store.dispatch('players/getPlayersByPage', this.current);
                }
            },
        },
    }
</script>

<style lang="scss">
    @import "../styles/_variables.scss";
    @import "../../node_modules/bulma/sass/utilities/_all.sass";
    @import "../../node_modules/bulma/sass/utilities/mixins.sass";

    @import "../../node_modules/bulma";
    @import "../../node_modules/buefy/src/scss/buefy.scss";
</style>
