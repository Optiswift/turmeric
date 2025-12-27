SRC_DIR="./lib/js/src"
PACKAGE_JSON="./lib/js/package.json"

clean_package_json() {
  jq '. += {"exports": {}, "typesVersions": {"*": {}}}' $PACKAGE_JSON > temp.json && mv temp.json $PACKAGE_JSON
}

update_package_json() {
  jq --arg project "$1" \
     --arg import_path "./src/$1/index.js" \
     --arg types_path "./src/$1/index.d.ts" \
     --arg types_version_path "src/$1/index.d.ts" \
    '.exports += {("./" + $project): {"import": $import_path, "require": $import_path, "types": $types_path}} |
     .typesVersions["*"] += {($project): [$types_version_path]}' \
    $PACKAGE_JSON > temp.json && mv temp.json $PACKAGE_JSON
}

generate_index_js() {
  {
    echo '"use strict";'
    echo 'Object.defineProperty(exports, "__esModule", { value: true });'
    for file in $1/*_pb.js; do
      echo "Object.assign(exports, require('./$(basename $file .js)'));"
    done
  } > "$1/index.js"
  chmod +x "$1/index.js"
}

generate_index_dts() {
  {
    for file in $1/*_pb.d.ts; do
      echo "export * from './$(basename $file .d.ts)';"
    done
  } > "$1/index.d.ts"
}

clean_package_json

task proto:generate:js

for project in $(ls ./proto | grep -v "^common$"); do
  PROJECT_DIR="$SRC_DIR/$project"
  if [ -d "$PROJECT_DIR" ]; then
    generate_index_js "$PROJECT_DIR"
    generate_index_dts "$PROJECT_DIR"
    update_package_json "$project"
  fi
done
