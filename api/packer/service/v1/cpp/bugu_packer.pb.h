// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: bugu_packer.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_bugu_5fpacker_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_bugu_5fpacker_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3019000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3019004 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_bugu_5fpacker_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_bugu_5fpacker_2eproto {
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::AuxiliaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::ParseTable schema[2]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::PROTOBUF_NAMESPACE_ID::internal::FieldMetadata field_metadata[];
  static const ::PROTOBUF_NAMESPACE_ID::internal::SerializationTable serialization_table[];
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_bugu_5fpacker_2eproto;
namespace bugu_packer {
namespace service {
namespace v1 {
class PackerReply;
struct PackerReplyDefaultTypeInternal;
extern PackerReplyDefaultTypeInternal _PackerReply_default_instance_;
class PackerRequest;
struct PackerRequestDefaultTypeInternal;
extern PackerRequestDefaultTypeInternal _PackerRequest_default_instance_;
}  // namespace v1
}  // namespace service
}  // namespace bugu_packer
PROTOBUF_NAMESPACE_OPEN
template<> ::bugu_packer::service::v1::PackerReply* Arena::CreateMaybeMessage<::bugu_packer::service::v1::PackerReply>(Arena*);
template<> ::bugu_packer::service::v1::PackerRequest* Arena::CreateMaybeMessage<::bugu_packer::service::v1::PackerRequest>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace bugu_packer {
namespace service {
namespace v1 {

// ===================================================================

class PackerRequest final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:bugu_packer.service.v1.PackerRequest) */ {
 public:
  inline PackerRequest() : PackerRequest(nullptr) {}
  ~PackerRequest() override;
  explicit constexpr PackerRequest(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  PackerRequest(const PackerRequest& from);
  PackerRequest(PackerRequest&& from) noexcept
    : PackerRequest() {
    *this = ::std::move(from);
  }

  inline PackerRequest& operator=(const PackerRequest& from) {
    CopyFrom(from);
    return *this;
  }
  inline PackerRequest& operator=(PackerRequest&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const PackerRequest& default_instance() {
    return *internal_default_instance();
  }
  static inline const PackerRequest* internal_default_instance() {
    return reinterpret_cast<const PackerRequest*>(
               &_PackerRequest_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(PackerRequest& a, PackerRequest& b) {
    a.Swap(&b);
  }
  inline void Swap(PackerRequest* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(PackerRequest* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  PackerRequest* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<PackerRequest>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const PackerRequest& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom(const PackerRequest& from);
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message* to, const ::PROTOBUF_NAMESPACE_ID::Message& from);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(PackerRequest* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "bugu_packer.service.v1.PackerRequest";
  }
  protected:
  explicit PackerRequest(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kDataFieldNumber = 1,
    kSizeFieldNumber = 2,
  };
  // repeated bytes data = 1;
  int data_size() const;
  private:
  int _internal_data_size() const;
  public:
  void clear_data();
  const std::string& data(int index) const;
  std::string* mutable_data(int index);
  void set_data(int index, const std::string& value);
  void set_data(int index, std::string&& value);
  void set_data(int index, const char* value);
  void set_data(int index, const void* value, size_t size);
  std::string* add_data();
  void add_data(const std::string& value);
  void add_data(std::string&& value);
  void add_data(const char* value);
  void add_data(const void* value, size_t size);
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>& data() const;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>* mutable_data();
  private:
  const std::string& _internal_data(int index) const;
  std::string* _internal_add_data();
  public:

  // uint32 size = 2;
  void clear_size();
  uint32_t size() const;
  void set_size(uint32_t value);
  private:
  uint32_t _internal_size() const;
  void _internal_set_size(uint32_t value);
  public:

  // @@protoc_insertion_point(class_scope:bugu_packer.service.v1.PackerRequest)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string> data_;
  uint32_t size_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_bugu_5fpacker_2eproto;
};
// -------------------------------------------------------------------

class PackerReply final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:bugu_packer.service.v1.PackerReply) */ {
 public:
  inline PackerReply() : PackerReply(nullptr) {}
  ~PackerReply() override;
  explicit constexpr PackerReply(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  PackerReply(const PackerReply& from);
  PackerReply(PackerReply&& from) noexcept
    : PackerReply() {
    *this = ::std::move(from);
  }

  inline PackerReply& operator=(const PackerReply& from) {
    CopyFrom(from);
    return *this;
  }
  inline PackerReply& operator=(PackerReply&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const PackerReply& default_instance() {
    return *internal_default_instance();
  }
  static inline const PackerReply* internal_default_instance() {
    return reinterpret_cast<const PackerReply*>(
               &_PackerReply_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    1;

  friend void swap(PackerReply& a, PackerReply& b) {
    a.Swap(&b);
  }
  inline void Swap(PackerReply* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(PackerReply* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  PackerReply* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<PackerReply>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const PackerReply& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom(const PackerReply& from);
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message* to, const ::PROTOBUF_NAMESPACE_ID::Message& from);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(PackerReply* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "bugu_packer.service.v1.PackerReply";
  }
  protected:
  explicit PackerReply(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  private:
  static void ArenaDtor(void* object);
  inline void RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena* arena);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  enum : int {
    kDataFieldNumber = 1,
    kSizeFieldNumber = 2,
  };
  // repeated bytes data = 1;
  int data_size() const;
  private:
  int _internal_data_size() const;
  public:
  void clear_data();
  const std::string& data(int index) const;
  std::string* mutable_data(int index);
  void set_data(int index, const std::string& value);
  void set_data(int index, std::string&& value);
  void set_data(int index, const char* value);
  void set_data(int index, const void* value, size_t size);
  std::string* add_data();
  void add_data(const std::string& value);
  void add_data(std::string&& value);
  void add_data(const char* value);
  void add_data(const void* value, size_t size);
  const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>& data() const;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>* mutable_data();
  private:
  const std::string& _internal_data(int index) const;
  std::string* _internal_add_data();
  public:

  // uint32 size = 2;
  void clear_size();
  uint32_t size() const;
  void set_size(uint32_t value);
  private:
  uint32_t _internal_size() const;
  void _internal_set_size(uint32_t value);
  public:

  // @@protoc_insertion_point(class_scope:bugu_packer.service.v1.PackerReply)
 private:
  class _Internal;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string> data_;
  uint32_t size_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_bugu_5fpacker_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// PackerRequest

// repeated bytes data = 1;
inline int PackerRequest::_internal_data_size() const {
  return data_.size();
}
inline int PackerRequest::data_size() const {
  return _internal_data_size();
}
inline void PackerRequest::clear_data() {
  data_.Clear();
}
inline std::string* PackerRequest::add_data() {
  std::string* _s = _internal_add_data();
  // @@protoc_insertion_point(field_add_mutable:bugu_packer.service.v1.PackerRequest.data)
  return _s;
}
inline const std::string& PackerRequest::_internal_data(int index) const {
  return data_.Get(index);
}
inline const std::string& PackerRequest::data(int index) const {
  // @@protoc_insertion_point(field_get:bugu_packer.service.v1.PackerRequest.data)
  return _internal_data(index);
}
inline std::string* PackerRequest::mutable_data(int index) {
  // @@protoc_insertion_point(field_mutable:bugu_packer.service.v1.PackerRequest.data)
  return data_.Mutable(index);
}
inline void PackerRequest::set_data(int index, const std::string& value) {
  data_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::set_data(int index, std::string&& value) {
  data_.Mutable(index)->assign(std::move(value));
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::set_data(int index, const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  data_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set_char:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::set_data(int index, const void* value, size_t size) {
  data_.Mutable(index)->assign(
    reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:bugu_packer.service.v1.PackerRequest.data)
}
inline std::string* PackerRequest::_internal_add_data() {
  return data_.Add();
}
inline void PackerRequest::add_data(const std::string& value) {
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::add_data(std::string&& value) {
  data_.Add(std::move(value));
  // @@protoc_insertion_point(field_add:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::add_data(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add_char:bugu_packer.service.v1.PackerRequest.data)
}
inline void PackerRequest::add_data(const void* value, size_t size) {
  data_.Add()->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_add_pointer:bugu_packer.service.v1.PackerRequest.data)
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>&
PackerRequest::data() const {
  // @@protoc_insertion_point(field_list:bugu_packer.service.v1.PackerRequest.data)
  return data_;
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>*
PackerRequest::mutable_data() {
  // @@protoc_insertion_point(field_mutable_list:bugu_packer.service.v1.PackerRequest.data)
  return &data_;
}

// uint32 size = 2;
inline void PackerRequest::clear_size() {
  size_ = 0u;
}
inline uint32_t PackerRequest::_internal_size() const {
  return size_;
}
inline uint32_t PackerRequest::size() const {
  // @@protoc_insertion_point(field_get:bugu_packer.service.v1.PackerRequest.size)
  return _internal_size();
}
inline void PackerRequest::_internal_set_size(uint32_t value) {
  
  size_ = value;
}
inline void PackerRequest::set_size(uint32_t value) {
  _internal_set_size(value);
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerRequest.size)
}

// -------------------------------------------------------------------

// PackerReply

// repeated bytes data = 1;
inline int PackerReply::_internal_data_size() const {
  return data_.size();
}
inline int PackerReply::data_size() const {
  return _internal_data_size();
}
inline void PackerReply::clear_data() {
  data_.Clear();
}
inline std::string* PackerReply::add_data() {
  std::string* _s = _internal_add_data();
  // @@protoc_insertion_point(field_add_mutable:bugu_packer.service.v1.PackerReply.data)
  return _s;
}
inline const std::string& PackerReply::_internal_data(int index) const {
  return data_.Get(index);
}
inline const std::string& PackerReply::data(int index) const {
  // @@protoc_insertion_point(field_get:bugu_packer.service.v1.PackerReply.data)
  return _internal_data(index);
}
inline std::string* PackerReply::mutable_data(int index) {
  // @@protoc_insertion_point(field_mutable:bugu_packer.service.v1.PackerReply.data)
  return data_.Mutable(index);
}
inline void PackerReply::set_data(int index, const std::string& value) {
  data_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::set_data(int index, std::string&& value) {
  data_.Mutable(index)->assign(std::move(value));
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::set_data(int index, const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  data_.Mutable(index)->assign(value);
  // @@protoc_insertion_point(field_set_char:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::set_data(int index, const void* value, size_t size) {
  data_.Mutable(index)->assign(
    reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_set_pointer:bugu_packer.service.v1.PackerReply.data)
}
inline std::string* PackerReply::_internal_add_data() {
  return data_.Add();
}
inline void PackerReply::add_data(const std::string& value) {
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::add_data(std::string&& value) {
  data_.Add(std::move(value));
  // @@protoc_insertion_point(field_add:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::add_data(const char* value) {
  GOOGLE_DCHECK(value != nullptr);
  data_.Add()->assign(value);
  // @@protoc_insertion_point(field_add_char:bugu_packer.service.v1.PackerReply.data)
}
inline void PackerReply::add_data(const void* value, size_t size) {
  data_.Add()->assign(reinterpret_cast<const char*>(value), size);
  // @@protoc_insertion_point(field_add_pointer:bugu_packer.service.v1.PackerReply.data)
}
inline const ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>&
PackerReply::data() const {
  // @@protoc_insertion_point(field_list:bugu_packer.service.v1.PackerReply.data)
  return data_;
}
inline ::PROTOBUF_NAMESPACE_ID::RepeatedPtrField<std::string>*
PackerReply::mutable_data() {
  // @@protoc_insertion_point(field_mutable_list:bugu_packer.service.v1.PackerReply.data)
  return &data_;
}

// uint32 size = 2;
inline void PackerReply::clear_size() {
  size_ = 0u;
}
inline uint32_t PackerReply::_internal_size() const {
  return size_;
}
inline uint32_t PackerReply::size() const {
  // @@protoc_insertion_point(field_get:bugu_packer.service.v1.PackerReply.size)
  return _internal_size();
}
inline void PackerReply::_internal_set_size(uint32_t value) {
  
  size_ = value;
}
inline void PackerReply::set_size(uint32_t value) {
  _internal_set_size(value);
  // @@protoc_insertion_point(field_set:bugu_packer.service.v1.PackerReply.size)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__
// -------------------------------------------------------------------


// @@protoc_insertion_point(namespace_scope)

}  // namespace v1
}  // namespace service
}  // namespace bugu_packer

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_bugu_5fpacker_2eproto
