<template>
  <div class="tinymce">
    <editor id="tinymce" v-model="a" :init="init"></editor>
  </div>
</template>

<script>
import tinymce from "tinymce/tinymce";
import "tinymce/themes/silver/theme";
import Editor from "@tinymce/tinymce-vue";
import "tinymce/plugins/image";
import "tinymce/plugins/link";
import "tinymce/plugins/code";
import "tinymce/plugins/table";
import "tinymce/plugins/lists";
import "tinymce/plugins/wordcount";
import "tinymce/icons/default/icons";
export default {
  name: "tinymce",
  props: {
    value: String
  },
  data() {
    return {
      a: this.value,
      init: {
        language_url: "/tinymce/zh_CN.js",
        language: "zh_CN",
        skin_url: "/tinymce/skins/ui/oxide",
        height: 300,
        plugins: "link lists image code table wordcount ",
        toolbar: "undo redo | styleselect | bold italic | link image",
        branding: false,
        init_instance_callback: (editor) => {
          editor.on("input", e => {
            this.$emit("input", e.target.innerHTML);
          });
          editor.on("change", e => {
            this.$emit('input', e.level.content)
          });
        }
      }
    };
  },
  watch:{
    'value': function(newVal){
      this.a = newVal
    }
  },
  mounted() {
    tinymce.init({});
  },
  components: { Editor }
};
</script>