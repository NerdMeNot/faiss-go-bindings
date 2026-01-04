/**
 * FAISS Go Extensions - C API Header
 *
 * This header declares additional C API functions for faiss-go that are either:
 * 1. Missing from the standard FAISS C API
 * 2. Have ABI issues in the standard C API
 *
 * Copyright (c) 2024 faiss-go contributors
 * Licensed under MIT License
 */

#ifndef FAISS_GO_EXT_H
#define FAISS_GO_EXT_H

#include <stdint.h>
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

/* Opaque pointer types - match FAISS C API */
typedef void* FaissIndex;
typedef void* FaissIndexBinary;
typedef void* FaissRangeSearchResult;
typedef void* FaissVectorTransform;

/* ============================================================
 * Index Assign Extension
 * ============================================================ */

/**
 * Assign vectors to their nearest neighbors (cluster assignment for IVF).
 *
 * @param index   The index
 * @param n       Number of vectors
 * @param x       Input vectors (n * d floats)
 * @param labels  Output labels (n * k int64_t)
 * @param k       Number of nearest neighbors to find
 * @return 0 on success, -1 on error
 */
int faiss_Index_assign_ext(FaissIndex index, int64_t n, const float* x, int64_t* labels, int64_t k);

/* ============================================================
 * Range Search Result Extensions
 * ============================================================ */

/**
 * Get the distances array from a RangeSearchResult.
 *
 * @param result    The RangeSearchResult pointer
 * @param distances Output pointer to the distances array
 * @return 0 on success, -1 on error
 */
int faiss_RangeSearchResult_distances(FaissRangeSearchResult result, float** distances);

/**
 * Get all arrays from a RangeSearchResult at once.
 *
 * @param result    The RangeSearchResult pointer
 * @param lims      Output: array of size nq+1 with result offsets
 * @param labels    Output: array of result labels
 * @param distances Output: array of result distances
 * @return 0 on success, -1 on error
 */
int faiss_RangeSearchResult_get(FaissRangeSearchResult result, int64_t** lims, int64_t** labels, float** distances);

/* ============================================================
 * Binary Index Constructor
 * ============================================================ */

/**
 * Create a new IndexBinaryFlat.
 *
 * @param p_index Output pointer to the new index
 * @param d       Dimension of binary vectors (in bits, must be multiple of 8)
 * @return 0 on success, -1 on error
 */
int faiss_IndexBinaryFlat_new(FaissIndexBinary* p_index, int64_t d);

/* ============================================================
 * HNSW Index Extensions (property accessors)
 * ============================================================ */

/**
 * Set efConstruction parameter for HNSW index.
 */
int faiss_IndexHNSW_set_efConstruction(FaissIndex index, int ef);

/**
 * Set efSearch parameter for HNSW index.
 */
int faiss_IndexHNSW_set_efSearch(FaissIndex index, int ef);

/**
 * Get efConstruction parameter from HNSW index.
 */
int faiss_IndexHNSW_get_efConstruction(FaissIndex index, int* ef);

/**
 * Get efSearch parameter from HNSW index.
 */
int faiss_IndexHNSW_get_efSearch(FaissIndex index, int* ef);

/* ============================================================
 * VectorTransform Extensions
 * ============================================================ */

/**
 * Train the VectorTransform.
 *
 * @param vt The VectorTransform
 * @param n  Number of vectors
 * @param x  Input vectors (n * d_in floats)
 * @return 0 on success, -1 on error
 */
int faiss_VectorTransform_train_ext(FaissVectorTransform vt, int64_t n, const float* x);

/**
 * Check if VectorTransform is trained.
 *
 * @param vt      The VectorTransform
 * @param trained Output: 1 if trained, 0 if not
 * @return 0 on success, -1 on error
 */
int faiss_VectorTransform_is_trained_ext(FaissVectorTransform vt, int* trained);

/**
 * Apply VectorTransform without allocating output.
 *
 * @param vt The VectorTransform
 * @param n  Number of vectors
 * @param x  Input vectors (n * d_in floats)
 * @param xt Output vectors (n * d_out floats, must be pre-allocated)
 * @return 0 on success, -1 on error
 */
int faiss_VectorTransform_apply_noalloc_ext(FaissVectorTransform vt, int64_t n, const float* x, float* xt);

/**
 * Reverse transform vectors.
 *
 * @param vt The VectorTransform
 * @param n  Number of vectors
 * @param xt Transformed vectors (n * d_out floats)
 * @param x  Output original vectors (n * d_in floats, must be pre-allocated)
 * @return 0 on success, -1 on error
 */
int faiss_VectorTransform_reverse_transform_ext(FaissVectorTransform vt, int64_t n, const float* xt, float* x);

#ifdef __cplusplus
}
#endif

#endif /* FAISS_GO_EXT_H */
