SRC_DIR="./lib/js/src"
PACKAGE_JSON="./lib/js/package.json"

clean_src_directory() {
  echo "Cleaning the src directory: $SRC_DIR"
  rm -rf "$SRC_DIR"/*
  jq '. += {"exports": {}, "typesVersions": {"*": {}}}' $PACKAGE_JSON > temp.json && mv temp.json $PACKAGE_JSON
}

# Update the exports field in package.json
update_package_json() {
  local project=$1
  local exports_import_path="./src/$project/index.js"
  local exports_types_path="./src/$project/index.d.ts"
  local types_version_path="src/$project/index.d.ts"

  # Add or update the export path for the project with both js and ts
  jq --arg project "$project" --arg import_path "$exports_import_path" --arg types_path "$exports_types_path" \
    '.exports += {("./" + $project): {"import": $import_path, "require": $import_path, "types": $types_path}}' \
    $PACKAGE_JSON > temp.json && mv temp.json $PACKAGE_JSON

  # Add or update the typesVersions path for the project
  jq --arg project "$project" --arg types_version_path "$types_version_path" \
    '.typesVersions["*"] += {($project): [$types_version_path]}' \
    $PACKAGE_JSON > temp.json && mv temp.json $PACKAGE_JSON
}

# Generate index.js for a project
generate_index_js() {
  local project_dir=$1

  echo "Generating index.js for $(basename $project_dir)"

  {
    echo "\"use strict\";"
    echo "Object.defineProperty(exports, \"__esModule\", { value: true });"
    for file in $(ls $project_dir/*_pb.js); do
      file_basename=$(basename $file .js)
      echo "Object.assign(exports, require('./$file_basename'));"
    done
  } > "$project_dir/index.js"

  chmod +x "$project_dir/index.js"
}


# Generate index.d.ts for a project
generate_index_dts() {
  local project_dir=$1

  echo "Generating index.ts for $(basename $project_dir)"

  {
    for file in $(ls $project_dir/*_pb.d.ts); do
      file_basename=$(basename $file .d.ts)
      echo "export * from './$file_basename';"
    done
  } > "$project_dir/index.d.ts"
}

clean_src_directory

for project in $(ls ./proto | grep -v "^common$"); \
do
  DEST="$SRC_DIR/$project"
  PROJECT=$project

  echo "Building proto for $project"
  mkdir -p $DEST

  grpc_tools_node_protoc \
    --js_out=import_style=commonjs,binary:$DEST \
    --grpc_out=grpc_js:$DEST \
    --ts_out=grpc_js:$DEST \
    -I ./proto/$project -I ./proto/common ./proto/$project/*.proto;

  generate_index_js "$DEST"
  generate_index_dts "$DEST"

  # Update package.json exports
  update_package_json "$project"

  echo "Completed proto for $project"
done
