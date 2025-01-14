<template>
  <v-row class="px-3 py-2">
    <v-col cols="3">
      <v-text-field
          label="Address"
          v-model="locomotive.address"
          variant="outlined"
          readonly
          density="compact"
      ></v-text-field>
    </v-col>
    <v-col cols="6">
      <v-text-field
          label="Name"
          v-model="locomotive.name"
          :hint="locomotive.name"
          variant="outlined"
          readonly
          density="compact"
      ></v-text-field>
    </v-col>
    <v-col cols="3">
      <v-text-field
          max="31"
          min="0"
          step="1"
          label="Speed"
          v-model="locomotive.speed"
          variant="outlined"
          type="number"
          :rules="rules"
          density="compact"
          @change="(e) => update()"
      ></v-text-field>
    </v-col>
  </v-row>

  <v-row class="px-3">
    <v-col cols="12">
      <v-expansion-panels multiple="true" mandatory="true" variant="accordion" color="deep-purple-lighten-2" v-model="panel">
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-railroad-light</v-icon>
            Train Motion
          </template>
          <v-expansion-panel-text>
            <v-row class="py-4">
              <v-col cols="8">
                <v-radio-group
                    v-model="locomotive.direction"
                    column
                    @change="(e) => update()"
                >
                  <v-radio
                      label="Forward >>"
                      :true-value="1"
                  ></v-radio>
                  <v-radio
                      label="<< Backward"
                      :true-value="0"
                  ></v-radio>
                </v-radio-group>
                <v-switch
                    :label="`Headlights: ${locomotive.fl === true ? 'On' : 'Off'}`"
                    :true-value="true"
                    :false-value="false"
                    color="amber-lighten-1"
                    v-model="locomotive.fl"
                    @change="(e) => update()"
                ></v-switch>
              </v-col>
              <v-col cols="4" class="text-center">
                <v-slider
                    justify-center
                    direction="vertical"
                    height="270"
                    max="31"
                    min="0"
                    step="1"
                    :focused="false"
                    :track-color="`rgb(${locomotive.speed * 10} 100 50)`"
                    :thumb-color="`rgb(${locomotive.speed * 10} 100 50)`"
                    :thumb-size="48"
                    :thumb-label="true"
                    v-model="locomotive.speed"
                    @update:model-value="(e) => {processSlider()}"
                ></v-slider>
                <v-btn
                    :disabled="locomotive.speed === 0 || !locomotive.speed"
                    variant="outlined"
                    color="red-darken-1"
                    @click="
            () => {
              locomotive.speed = 0;
              update();
            }
          "
                >
                  STOP
                </v-btn>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-checkbox-multiple-marked-outline</v-icon>
            F0 ... F4
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in 5" :key="i">
                <v-checkbox
                    :label="`F${i-1}`"
                    :true-value="true"
                    v-model="locomotive[i-1===0?'fl':'f'+(i-1)]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-checkbox>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-checkbox-multiple-marked-outline</v-icon>
            F5 ... F8
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in 4" :key="i">
                <v-checkbox
                    :label="`F${i+4}`"
                    :true-value="true"
                    v-model="locomotive['f'+(i+4)]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-checkbox>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-checkbox-multiple-marked-outline</v-icon>
            F9 ... F12
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in 4" :key="i">
                <v-checkbox
                    :label="`F${i+8}`"
                    :true-value="true"
                    v-model="locomotive['f'+(i+8)]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-checkbox>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-checkbox-multiple-marked-outline</v-icon>
            F13 ... F20
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in 8" :key="i">
                <v-checkbox
                    :label="`F${i+12}`"
                    :true-value="true"
                    v-model="locomotive['f'+(i+12)]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-checkbox>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
        <v-expansion-panel>
          <template v-slot:title>
            <v-icon class="mr-2">mdi-checkbox-multiple-marked-outline</v-icon>
            F21 ... F28
          </template>
          <v-expansion-panel-text>
            <v-row>
              <v-col cols="4" v-for="i in 8" :key="i">
                <v-checkbox
                    :label="`F${i+20}`"
                    :true-value="true"
                    v-model="locomotive['f'+(i+20)]"
                    color="success"
                    density="comfortable"
                    @change="(e) => update()"
                ></v-checkbox>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-col>
  </v-row>
</template>

<script setup>
import {computed, defineProps, ref} from "vue";
import {useControllerStore} from "../store/modules/controller";

const props = defineProps(['name'])

const store = useControllerStore()
const locomotive = computed(() => store.locomotives[props.name])

const rules = [
  (v) => v <= 31 || "Max speed is 31",
  (v) => v >= 0 || "Min speed is 0",
]

const funcNumber = 28
const targetSpeed = ref(0)
const panel = ref([0])

function update() {
  store.sendDataToServer()
}

function debounce(func, timeout = 500) {
  let timer;
  return (...args) => {
    clearTimeout(timer);
    timer = setTimeout(() => {
      func.apply(this, args);
    }, timeout);
  };
}

const processSlider = debounce(() => update());
</script>