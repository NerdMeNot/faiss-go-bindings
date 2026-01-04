// Package bindings provides pre-built FAISS static libraries for Go.
//
// This package is meant to be imported with a blank identifier:
//
//	import _ "github.com/NerdMeNot/faiss-go-bindings"
//
// The CGO directives in this package configure the linker to use the
// pre-built FAISS libraries bundled with this module.
//
// Supported platforms:
//   - linux/amd64, linux/arm64
//   - darwin/amd64, darwin/arm64
//   - windows/amd64
//
// For other platforms, build faiss-go with -tags=faiss_use_system
// to link against a system-installed FAISS.
package bindings

/*
#cgo LDFLAGS: -lstdc++ -lm

#include <stdlib.h>
#include <stdint.h>

// FAISS C API type declarations
// These are forward declarations for the FAISS C API types.
// The actual implementations are in the pre-built static libraries.

#ifdef __cplusplus
extern "C" {
#endif

// Opaque pointer types
typedef void* FaissIndex;
typedef void* FaissIndexBinary;
typedef void* FaissVectorTransform;
typedef void* FaissKmeans;
typedef void* FaissIndexIVF;
typedef void* FaissIndexIVFFlat;
typedef void* FaissIndexRefineFlat;
typedef void* FaissIndexPreTransform;
typedef void* FaissIndexShards;
typedef void* FaissRangeSearchResult;
typedef void* FaissClustering;

// ==== Flat Index Functions ====
extern int faiss_IndexFlatL2_new_with(FaissIndex* p_index, int64_t d);
extern int faiss_IndexFlatIP_new_with(FaissIndex* p_index, int64_t d);

// ==== IVF Index Functions ====
extern FaissIndexIVF* faiss_IndexIVF_cast(FaissIndex index);
extern int faiss_IndexIVFFlat_new_with_metric(FaissIndexIVFFlat** p_index, FaissIndex quantizer, size_t d, size_t nlist, int metric_type);
extern void faiss_IndexIVF_set_nprobe(FaissIndexIVF* index, size_t nprobe);
extern size_t faiss_IndexIVF_nprobe(FaissIndexIVF* index);
extern void faiss_IndexIVF_set_own_fields(FaissIndexIVF* index, int own_fields);

// ==== Scalar Quantizer Index Functions ====
extern int faiss_IndexScalarQuantizer_new_with(FaissIndex* p_index, int64_t d, int qtype, int metric_type);
extern int faiss_IndexIVFScalarQuantizer_new_with_metric(FaissIndex* p_index, FaissIndex quantizer, int64_t d, int64_t nlist, int qtype, int metric_type, int encode_residual);
extern void faiss_IndexIVFScalarQuantizer_set_own_fields(FaissIndex index, int own_fields);

// ==== LSH Index Functions ====
extern int faiss_IndexLSH_new(FaissIndex** p_index, int64_t d, int nbits);
extern int faiss_IndexLSH_new_with_options(FaissIndex** p_index, int64_t d, int nbits, int rotate_data, int train_thresholds);

// ==== ID Map Functions ====
extern int faiss_IndexIDMap_new(FaissIndex* p_index, FaissIndex base_index);
extern void faiss_IndexIDMap_set_own_fields(FaissIndex index, int own_fields);
extern int faiss_IndexIDMap_add_with_ids(FaissIndex index, int64_t n, const float* x, const int64_t* ids);

// ==== Common Index Operations ====
extern int faiss_Index_add(FaissIndex index, int64_t n, const float* x);
extern int faiss_Index_add_with_ids(FaissIndex index, int64_t n, const float* x, const int64_t* ids);
extern int faiss_Index_search(FaissIndex index, int64_t n, const float* x, int64_t k, float* distances, int64_t* labels);
extern int faiss_Index_train(FaissIndex index, int64_t n, const float* x);
extern int faiss_Index_assign(FaissIndex* index, int64_t n, const float* x, int64_t* labels, int64_t k);
extern int faiss_Index_reconstruct(FaissIndex index, int64_t key, float* recons);
extern int faiss_Index_reconstruct_n(FaissIndex index, int64_t i0, int64_t ni, float* recons);
extern int faiss_Index_reset(FaissIndex index);
extern void faiss_Index_free(FaissIndex index);
extern int64_t faiss_Index_ntotal(FaissIndex index);
extern int faiss_Index_is_trained(FaissIndex index);
extern int faiss_Index_d(FaissIndex index);

// ==== Range Search ====
extern int faiss_RangeSearchResult_new(FaissRangeSearchResult* p_result, int64_t nq);
extern int faiss_Index_range_search(FaissIndex index, int64_t n, const float* x, float radius, FaissRangeSearchResult result);
extern int faiss_RangeSearchResult_nq(FaissRangeSearchResult result);
extern size_t faiss_RangeSearchResult_buffer_size(FaissRangeSearchResult result);
extern int64_t* faiss_RangeSearchResult_lims(FaissRangeSearchResult result);
extern int64_t* faiss_RangeSearchResult_labels(FaissRangeSearchResult result);
extern float* faiss_RangeSearchResult_distances(FaissRangeSearchResult result);
extern int faiss_RangeSearchResult_get(FaissRangeSearchResult result, int64_t** lims, int64_t** labels, float** distances);
extern void faiss_RangeSearchResult_free(FaissRangeSearchResult result);

// ==== Index Factory ====
extern int faiss_index_factory(FaissIndex* p_index, int d, const char* description, int metric_type);

// ==== Serialization Functions ====
extern int faiss_write_index_fname(const FaissIndex* idx, const char* fname);
extern int faiss_read_index_fname(const char* fname, int io_flags, FaissIndex** p_out);

// ==== Binary Index Functions ====
extern int faiss_IndexBinaryFlat_new(FaissIndexBinary* p_index, int64_t d);
extern int faiss_IndexBinary_add(FaissIndexBinary index, int64_t n, const uint8_t* x);
extern int faiss_IndexBinary_search(FaissIndexBinary index, int64_t n, const uint8_t* x, int64_t k, int32_t* distances, int64_t* labels);
extern int faiss_IndexBinary_train(FaissIndexBinary index, int64_t n, const uint8_t* x);
extern int faiss_IndexBinary_reset(FaissIndexBinary index);
extern int faiss_IndexBinary_ntotal(FaissIndexBinary index, int64_t* ntotal);
extern int faiss_IndexBinary_is_trained(FaissIndexBinary index, int* is_trained);
extern int faiss_IndexBinaryIVF_set_nprobe(FaissIndexBinary index, int64_t nprobe);
extern void faiss_IndexBinary_free(FaissIndexBinary index);

// ==== HNSW Property Accessors (from faiss_go_ext) ====
extern int faiss_IndexHNSW_set_efConstruction(FaissIndex index, int ef);
extern int faiss_IndexHNSW_set_efSearch(FaissIndex index, int ef);
extern int faiss_IndexHNSW_get_efConstruction(FaissIndex index, int* ef);
extern int faiss_IndexHNSW_get_efSearch(FaissIndex index, int* ef);

// ==== Index Assign Extension ====
extern int faiss_Index_assign_ext(FaissIndex index, int64_t n, const float* x, int64_t* labels, int64_t k);

// ==== Vector Transform Functions ====
extern int faiss_PCAMatrix_new_with(FaissVectorTransform* p_transform, int64_t d_in, int64_t d_out, float eigen_power, int random_rotation);
extern int faiss_OPQMatrix_new_with(FaissVectorTransform* p_transform, int d, int M, int d2);
extern int faiss_RandomRotationMatrix_new_with(FaissVectorTransform* p_transform, int64_t d_in, int64_t d_out);
extern int faiss_VectorTransform_train_ext(FaissVectorTransform vt, int64_t n, const float* x);
extern int faiss_VectorTransform_is_trained_ext(FaissVectorTransform vt, int* trained);
extern int faiss_VectorTransform_apply_noalloc_ext(FaissVectorTransform vt, int64_t n, const float* x, float* xt);
extern int faiss_VectorTransform_reverse_transform_ext(FaissVectorTransform vt, int64_t n, const float* xt, float* x);
extern void faiss_VectorTransform_free(FaissVectorTransform vt);

// ==== Clustering Functions ====
extern int faiss_Clustering_new(FaissClustering* p_clustering, int d, int k);
extern int faiss_Clustering_train(FaissClustering clustering, int64_t n, const float* x, FaissIndex index);
extern void faiss_Clustering_centroids(FaissClustering* clustering, float** centroids, size_t* size);
extern void faiss_Clustering_free(FaissClustering clustering);
extern int faiss_kmeans_clustering(size_t d, size_t n, size_t k, const float* x, float* centroids, float* q_error);

// ==== Composite Index Functions ====
extern int faiss_IndexRefineFlat_new(FaissIndex* p_index, FaissIndex base_index);
extern void faiss_IndexRefineFlat_set_k_factor(FaissIndexRefineFlat index, float k_factor);
extern void faiss_IndexRefineFlat_set_own_fields(FaissIndex index, int own_fields);
extern int faiss_IndexPreTransform_new_with_transform(FaissIndexPreTransform** p_index, FaissVectorTransform* ltrans, FaissIndex* index);
extern void faiss_IndexPreTransform_set_own_fields(FaissIndex index, int own_fields);
extern int faiss_IndexShards_new(FaissIndexShards** p_index, int64_t d);
extern int faiss_IndexShards_add_shard(FaissIndexShards* index, FaissIndex* shard);

#ifdef __cplusplus
}
#endif
*/
import "C"
import "unsafe"

// CreateTestIndex creates a simple FlatL2 index for testing.
// Returns the index pointer or an error code.
func CreateTestIndex(dim int) (uintptr, int) {
	var idx C.FaissIndex
	ret := C.faiss_IndexFlatL2_new_with(&idx, C.int64_t(dim))
	if ret != 0 {
		return 0, int(ret)
	}
	return uintptr(unsafe.Pointer(idx)), 0
}

// GetIndexDimension returns the dimension of an index.
func GetIndexDimension(ptr uintptr) int {
	idx := C.FaissIndex(unsafe.Pointer(ptr))
	return int(C.faiss_Index_d(idx))
}

// GetIndexNtotal returns the number of vectors in an index.
func GetIndexNtotal(ptr uintptr) int64 {
	idx := C.FaissIndex(unsafe.Pointer(ptr))
	return int64(C.faiss_Index_ntotal(idx))
}

// FreeIndex frees an index.
func FreeIndex(ptr uintptr) {
	idx := C.FaissIndex(unsafe.Pointer(ptr))
	C.faiss_Index_free(idx)
}
