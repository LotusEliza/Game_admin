<template>
  <div id="app" class="bg-img">
    <div v-cloak>
      <app-header>
        <app-search class="is-hidden-touch"></app-search>
      </app-header>
      <transition name="fade" mode="out-in" v-on:after-enter="afterEnter" appear>
        <router-view class="bg-img"></router-view>
      </transition>
    </div>
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Search from './components/SearchPlayer.vue'

export default {
  name: 'app',
  components: {
    'app-header': Header,
    'app-search': Search,
  },
  data(){
    return{
      hello:'hello title',
    }
  }
}
</script>

<style>

  [v-cloak] { display: none; }

  .p-3 {
    padding: 20px 50px 0px 50px !important;
  }
  .pd-3 {
    padding: 0px 50px 50px 50px !important;
  }
  .pt-2 {
    padding-top: 20px !important;
  }
  .p-4 {
    padding: 10px 200px 0px 200px !important;
  }
  .m-3 {
    margin: 0px 5px 0px 5px !important;
  }
  .pb-3 {
    padding-bottom: 10px !important;
  }

  /*.container{*/
  /*  margin: 0px !important;*/
  /*  padding: 0px !important;*/
  /*}*/


  /*html {*/
  /*  background: url('https://www.wallpaperflare.com/static/472/195/776/planet-sun-digital-art-space-wallpaper.jpg') no-repeat center center fixed;*/
  /*  -webkit-background-size: cover;*/
  /*  -moz-background-size: cover;*/
  /*  -o-background-size: cover;*/
  /*  background-size: cover;*/
  /*}*/


  .bg-img {
    background-image: url("https://www.wallpaperflare.com/static/472/195/776/planet-sun-digital-art-space-wallpaper.jpg") ;
    background-position: center center;
    background-repeat:  no-repeat;
    background-attachment: fixed;
    background-size:  cover;
    background-color: #999;

  }
  .tile.is-child.box {
    background-color: rgba(248, 229, 190, 0.6) !important;
  }

  body {
    font-size: 14px !important;
    background-color: grey;
  }

  .title {
    /*color: #363636;*/
    font-size: 1.5rem !important;
    line-height: 0.5 !important;
    /*font-weight: 600;*/
    /*line-height: 1.125;*/
  }
  .subtitle {
    font-size: 1rem !important;
    line-height: 0.9 !important;
  }

  /*TRANSITION*/

  .fade-enter-active, .fade-leave-active {
    transition: opacity 0.2s
  }

  .fade-enter, .fade-leave-active {
    opacity: 0
  }
</style>
