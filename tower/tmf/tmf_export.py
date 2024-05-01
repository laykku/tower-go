bl_info = {
    "name": "TMFExport",
    "description": "Exports scene to a Tower Engine mesh data format",
    "author": "laykku",
    "version": (1, 0),
    "blender": (4,0,0),
    "location": "File > Export",
    "category": "Import-Export"
}

import bpy
import struct

def export_objects(filepath):
    for obj in bpy.data.objects:    
        if obj.type == 'MESH':
            obj_name = None
            vertices = []
            uv0 = []
            indices = []

            # triangulate

            bpy.context.view_layer.objects.active = obj
            bpy.ops.object.mode_set(mode='EDIT')
            bpy.ops.mesh.select_all(action='SELECT')

            bpy.ops.mesh.quads_convert_to_tris()
            bpy.ops.object.mode_set(mode='OBJECT')

            # -----

            obj_name = obj.name.encode('utf-8')

            mesh = obj.data

            uv_layer = None
            if mesh.uv_layers.active:
                uv_layer = mesh.uv_layers.active.data

            for vertex in mesh.vertices:
                vertices.append(vertex.co)

            for polygon in mesh.polygons:
                for loop_index in polygon.loop_indices:
                    vertex_index = mesh.loops[loop_index].vertex_index
                    uv0.append(uv_layer[loop_index].uv)

            for polygon in mesh.polygons:
                for vertex_index in polygon.vertices:
                    indices.append(vertex_index)

            with open(filepath, "wb") as f:
                f.write(struct.pack('i', len(obj_name)))
                f.write(obj_name)
                
                f.write(struct.pack('i', len(vertices)))
                for vertex in vertices:
                    f.write(struct.pack('fff', *vertex))
                
                f.write(struct.pack('i', len(uv0)))
                for uv in uv0:
                    f.write(struct.pack('ff', *uv))

                f.write(struct.pack('i', len(indices)))
                for index in indices:
                    f.write(struct.pack('i', index))

            break # todo store mesh data in array with mesh name as key

def menu_func(self, context):
    self.layout.operator(TMFExport.bl_idname, text="Tower Mesh Format")

class TMFExport(bpy.types.Operator):
    bl_idname = "object.tmfexport"
    bl_label = "Export"
    bl_description = "Tower Engine mesh exporter"

    filepath: bpy.props.StringProperty(subtype="FILE_PATH")

    def execute(self, context):
        path = self.filepath
        if not path.endswith(".tmf"):
            path += ".tmf"
        export_objects(path)
        return {'FINISHED'}

    def invoke(self, context, event):
        context.window_manager.fileselect_add(self)
        return {'RUNNING_MODAL'}

def register():
    bpy.utils.register_class(TMFExport)
    bpy.types.TOPBAR_MT_file_export.append(menu_func)

def unregister():
    bpy.utils.unregister_class(TMFExport)
    bpy.type.TOPBAR_MT_file_export.remove(menu_func)
