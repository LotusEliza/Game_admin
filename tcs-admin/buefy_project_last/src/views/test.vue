<template>
    <div id="app" class="container">

        {{checkedRows}}
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
        <section>
            <b-table
                    :data="isEmpty ? [] : players"
                    :columns="columns"
                    :checked-rows.sync="checkedRows"
                    checkable
                    :checkbox-position="right"
                    :striped="isStriped"
                    :narrowed="isNarrowed"
                    :hoverable="isHoverable"
                    :loading="isLoading"
                    :mobile-cards="hasMobileCards"
            >
                <template slot="right">
                    Hello
                </template>
            </b-table>

        </section>

    </div>

</template>

<script>
    import { mapGetters } from 'vuex';
    export default {
        data() {
            return {

                isEmpty: false,
                isStriped: false,
                isNarrowed: false,
                isHoverable: true,
                isLoading: false,
                hasMobileCards: true,
                narrowed: true,
                checkedRows: [],

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

                columns: [
                    {
                        field: 'Tg',
                        label: 'Tg',
                        width: '90',
                        numeric: true,
                        searchable: true,
                    },
                    {
                        field: 'Name',
                        label: 'Name',
                        searchable: true,
                    },
                    {
                        field: 'Story',
                        label: 'Story',
                        searchable: true,
                    },
                    {
                        field: 'Faction',
                        label: 'Faction',
                        centered: true
                    },
                    {
                        field: 'Referrer',
                        label: 'Referrer',
                    },
                    {
                        field: 'Location',
                        label: 'Location',
                    },
                    {
                        field: 'PosX',
                        label: 'PosX',
                    },
                    {
                        field: 'PosY',
                        label: 'PosY',
                    },
                    {
                        field: 'HP',
                        label: 'HP',
                    },
                    {
                        field: 'EquipW',
                        label: 'EquipW',
                    },
                    {
                        field: 'EquipA',
                        label: 'EquipA',
                    },
                    {
                        field: 'EquipB',
                        label: 'EquipB',
                    },
                    {
                        field: 'EquipC',
                        label: 'EquipC',
                    },
                    {
                        field: 'Air',
                        label: 'Air',
                    },
                    {
                        field: 'Registered',
                        label: 'Registered',
                    },
                    {
                        field: 'LastActive',
                        label: 'LastActive',
                    },
                ]
            }
        },
        computed:{
            ...mapGetters([
                'total_players',
                'players',
            ]),
        },
        mounted(){
            this.$store.dispatch('getPlayersByPage', this.current);
        },
        watch: {
            current: {
                handler: function() {
                    this.$store.dispatch('getPlayersByPage', this.current);
                }
            },
        },
    }
    //
    // const app = new Vue(example)
    // app.$mount('#app')

</script>
