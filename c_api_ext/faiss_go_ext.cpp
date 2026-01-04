/**
 * FAISS Go Extensions - Custom C API wrappers for faiss-go
 *
 * This file provides additional C wrapper functions that are either:
 * 1. Missing from the standard FAISS C API
 * 2. Have ABI issues in the standard C API
 *
 * Compile with:
 *   g++ -c -fPIC -std=c++17 -I/path/to/faiss/include faiss_go_ext.cpp -o faiss_go_ext.o
 *   ar rcs libfaiss_go_ext.a faiss_go_ext.o
 *
 * Copyright (c) 2024 faiss-go contributors
 * Licensed under MIT License
 */

#include "faiss_go_ext.h"

#include <faiss/Index.h>
#include <faiss/IndexHNSW.h>
#include <faiss/IndexBinaryFlat.h>
#include <faiss/VectorTransform.h>
#include <faiss/impl/AuxIndexStructures.h>
#include <cstdint>
#include <cstring>

extern "C" {

// ============================================================
// Index Assign - Custom wrapper that works reliably
// ============================================================

int faiss_Index_assign_ext(FaissIndex index, int64_t n, const float* x, int64_t* labels, int64_t k) {
    try {
        if (!index) return -1;
        auto* idx = static_cast<faiss::Index*>(index);
        idx->assign(n, x, labels, k);
        return 0;
    } catch (...) {
        return -1;
    }
}

// ============================================================
// Range Search Result Extensions
// ============================================================

int faiss_RangeSearchResult_distances(FaissRangeSearchResult result, float** distances) {
    if (!result || !distances) return -1;
    auto* res = static_cast<faiss::RangeSearchResult*>(result);
    *distances = res->distances;
    return 0;
}

int faiss_RangeSearchResult_get(FaissRangeSearchResult result, int64_t** lims, int64_t** labels, float** distances) {
    if (!result) return -1;
    auto* res = static_cast<faiss::RangeSearchResult*>(result);
    if (lims) *lims = reinterpret_cast<int64_t*>(res->lims);
    if (labels) *labels = reinterpret_cast<int64_t*>(res->labels);
    if (distances) *distances = res->distances;
    return 0;
}

// ============================================================
// HNSW Property Accessors
// ============================================================

int faiss_IndexHNSW_set_efConstruction(FaissIndex index, int ef) {
    try {
        auto* hnsw = dynamic_cast<faiss::IndexHNSW*>(static_cast<faiss::Index*>(index));
        if (!hnsw) return -1;
        hnsw->hnsw.efConstruction = ef;
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_IndexHNSW_set_efSearch(FaissIndex index, int ef) {
    try {
        auto* hnsw = dynamic_cast<faiss::IndexHNSW*>(static_cast<faiss::Index*>(index));
        if (!hnsw) return -1;
        hnsw->hnsw.efSearch = ef;
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_IndexHNSW_get_efConstruction(FaissIndex index, int* ef) {
    try {
        auto* hnsw = dynamic_cast<faiss::IndexHNSW*>(static_cast<faiss::Index*>(index));
        if (!hnsw || !ef) return -1;
        *ef = hnsw->hnsw.efConstruction;
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_IndexHNSW_get_efSearch(FaissIndex index, int* ef) {
    try {
        auto* hnsw = dynamic_cast<faiss::IndexHNSW*>(static_cast<faiss::Index*>(index));
        if (!hnsw || !ef) return -1;
        *ef = hnsw->hnsw.efSearch;
        return 0;
    } catch (...) {
        return -1;
    }
}

// ============================================================
// Binary Index Constructors
// ============================================================

int faiss_IndexBinaryFlat_new(FaissIndexBinary* p_index, int64_t d) {
    try {
        *p_index = new faiss::IndexBinaryFlat(d);
        return 0;
    } catch (...) {
        return -1;
    }
}

// ============================================================
// VectorTransform Extensions - Custom wrappers for ABI safety
// ============================================================

int faiss_VectorTransform_train_ext(FaissVectorTransform vt, int64_t n, const float* x) {
    try {
        if (!vt) return -1;
        auto* transform = static_cast<faiss::VectorTransform*>(vt);
        transform->train(n, x);
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_VectorTransform_is_trained_ext(FaissVectorTransform vt, int* trained) {
    try {
        if (!vt || !trained) return -1;
        auto* transform = static_cast<faiss::VectorTransform*>(vt);
        *trained = transform->is_trained ? 1 : 0;
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_VectorTransform_apply_noalloc_ext(FaissVectorTransform vt, int64_t n, const float* x, float* xt) {
    try {
        if (!vt) return -1;
        auto* transform = static_cast<faiss::VectorTransform*>(vt);
        transform->apply_noalloc(n, x, xt);
        return 0;
    } catch (...) {
        return -1;
    }
}

int faiss_VectorTransform_reverse_transform_ext(FaissVectorTransform vt, int64_t n, const float* xt, float* x) {
    try {
        if (!vt) return -1;
        auto* transform = static_cast<faiss::VectorTransform*>(vt);
        transform->reverse_transform(n, xt, x);
        return 0;
    } catch (...) {
        return -1;
    }
}

} // extern "C"
