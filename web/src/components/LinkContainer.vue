<script setup lang="ts">
import { ref, toRefs, reactive, onMounted } from 'vue'

const props = defineProps<{ tag: string }>()

// get the properties
const { tag } = toRefs(props);

const show = ref(true)

// const linkList = reactive<any[]>([])
const linkList = ref<any[]>([])

onMounted(() => {
  fetch("/tag/"+tag.value )
  .then( r => r.json())
  .then( (arr:any[]) => {    
    arr.sort( (a:any,b:any) => {
      return a.create_at.localeCompare(b.create_at)
    })
    linkList.value = arr
  })
})

function bgURL(url:string){
  return `background-image: url("${url}/favicon.ico")`
}


// const linkList = async (params:string) => {
//   return fetch("/tag/"+params)
//           .then(r => r.json())
// }
</script>

<template>
  <div class="link-contianer">
    <h2 class="link-continer title" @click="$event => {show = !show}">{{ tag }}</h2>
      <!-- {{ show }}
      {{ linkList }} -->
      <div v-show="show">
        <a v-for="link in linkList" :key="link.id" 
          :href="'/go?id='+link.id"
          :style="bgURL(link.url)"
          class="card" 
        >
          <div class="container">
            <h3 class="title">{{ link.title }}</h3>
            <div class="description">{{ link.description }}</div>
          </div>
        </a>

      </div>
  </div>
  <br style="clear:both;">
</template>

<style >

.link-contianer {
  clear: both;
  margin: 1em;
}

h2 {
  cursor: pointer;
  background-color: rgba(255,255,255,0.5);
  padding: 0.5em;
  opacity: 1;
  text-shadow: 0.2rem 0rem 0.5rem white, -0.2rem 0rem 0.5rem white,
    0rem 0.2rem 0.5rem white, 0rem -0.2rem 0.5rem white;
}

h3 {
  margin: 0em;
  padding: 0em;
}

.continer * {
  padding: 0em 1em 0em 0em;
}


a {
  text-decoration: none;
}

a:hover {
  text-decoration: none;
}



.card .s-hover {
  display: none;
}
.card:hover .s-hover {
  font-size: medium;
  color: rgba(255, 96, 202, 0.6);
  display: inline-block;
}

.card {
  width: 400px;
  height: 85px;
  float: left;
  margin: 0.5em;
  padding: 1em;
  background-color: rgba(255, 255, 255, 0.5);

  background-repeat: no-repeat;
  background-size: contain;
  background-blend-mode: soft-light;
  
  overflow-y: hidden;
}

.card:hover {
  background-color: rgba(255, 255, 255, 0.8);
  cursor: pointer;
}


.s-describe {
  font-size: large;
  color: rgba(0, 0, 0, 0.8);
}

.description{
  margin-left: 2em;
}

.card>.container {
  font-size: large;
  color: rgba(0, 0, 0, 0.8);
  overflow-wrap: anywhere;
}


</style>
